package connection

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"reflect"
	"sync"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/utils"
	"ultipa-go-sdk/sdk/utils/logger"
)

type GraphClusterInfo struct {
	Graph         string
	Leader        *Connection
	Followers     []*Connection
	Algos         []*Connection
	LastAlgoIndex int //记录上次使用的 Task 节点索引
}

// handle all connections
type ConnectionPool struct {
	GraphMgr        *GraphManager // graph name : ClusterInfo
	Config          *configuration.UltipaConfig
	Connections     map[string]*Connection // Host : Connection
	RandomTick      int
	Actives         []*Connection
	LastActivesTime time.Time
	IsRaft          bool
	muActiveSafely  sync.Mutex
}

func NewConnectionPool(config *configuration.UltipaConfig) (*ConnectionPool, error) {

	if len(config.Hosts) < 1 {
		return nil, errors.New("Error Hosts can not by empty")
	}

	pool := &ConnectionPool{
		Config:      config,
		Connections: map[string]*Connection{},
		GraphMgr:    NewGraphManager(),
	}

	// Init Cluster Manager
	// Get Connections
	err := pool.CreateConnections()

	if err != nil {
		log.Println(err)
	}

	// Refresh Actives
	pool.RefreshActives()

	// Refresh global Cluster info
	err = pool.RefreshClusterInfo("global")

	if err != nil {
		log.Println(err)
	}

	return pool, err
}

func (pool *ConnectionPool) CreateConnections() error {
	var err error

	for _, host := range pool.Config.Hosts {
		conn, _ := NewConnection(host, pool.Config)
		pool.Connections[host] = conn
	}

	return err
}

func (pool *ConnectionPool) RefreshActivesWithSeconds(seconds int32) {
	pool.muActiveSafely.Lock()
	defer pool.muActiveSafely.Unlock()
	if time.Now().Sub(pool.LastActivesTime) <= 5*time.Second && len(pool.Connections) == len(pool.Actives) {
		// 避免频繁刷新
		return
	}
	defer func() {
		pool.LastActivesTime = time.Now()
	}()
	pool.Actives = []*Connection{}
	if seconds <= 0 {
		seconds = 3
	}
	var wg sync.WaitGroup
	for _, conn := range pool.Connections {
		wg.Add(1)
		go func(conn *Connection) {
			defer wg.Done()
			ctx, cancel, err := pool.NewContext(&configuration.RequestConfig{
				Timeout: seconds,
			})
			if err != nil {
				logger.PrintWarn(conn.Host + "failed - " + err.Error())
				conn.Active = ultipa.ServerStatus_DEAD
				return
			}
			defer cancel()

			resp, err := conn.GetControlClient().SayHello(ctx, &ultipa.HelloUltipaRequest{
				Name: "go sdk refresh",
			})

			if err != nil {
				logger.PrintWarn(conn.Host + "failed - " + err.Error())
				conn.Active = ultipa.ServerStatus_DEAD
				return
			}

			if resp.Status == nil || resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS {
				conn.Active = ultipa.ServerStatus_ALIVE
				pool.Actives = append(pool.Actives, conn)
			} else {
				logger.PrintWarn(conn.Host + "failed - " + resp.Status.Msg)
				conn.Active = ultipa.ServerStatus_DEAD
			}
		}(conn)
	}
	wg.Wait()
}

// 更新查看哪些连接还有效
func (pool *ConnectionPool) RefreshActives() {
	pool.RefreshActivesWithSeconds(6)
}
func (pool *ConnectionPool) ForceRefreshClusterInfo(graphName string) error {
	pool.GraphMgr.DeleteGraph(graphName)
	return pool.RefreshClusterInfo(graphName)
}

// sync cluster info from server
func (pool *ConnectionPool) RefreshClusterInfo(graphName string) error {
	err := pool.doRefreshClusterInfo(graphName)
	if err != nil && reflect.TypeOf(err).Elem().String() == "utils.LeaderNotYetElectedError" {
		//若是leader未选出的错误类型，再重试一次
		pool.RefreshActives()
		err = pool.doRefreshClusterInfo(graphName)
	}
	return err
}

func (pool *ConnectionPool) doRefreshClusterInfo(graphName string) error {
	var conn *Connection

	var err error

	activeConns := pool.Actives

	if len(pool.Actives) < 1 {
		return errors.New("no active Connection is found")
	}

	allIsNill := true
	for _, activeConn := range activeConns {
		if activeConn == nil {
			continue
		}
		allIsNill = false
		// 如果该图集暂无初始化时
		if pool.GraphMgr.GetLeader(graphName) == nil {
			conn = activeConn
		} else {
			// 已经初始化后
			conn = pool.GraphMgr.GetLeader(graphName)
		}
		logger.PrintInfo(fmt.Sprintf("refresh graph [%s] cluster info with connection to host [%s]", graphName, conn.Host))
		err = pool.resolveClusterInfo(graphName, conn)
		if err == nil {
			return nil
		}
	}
	if allIsNill {
		err = errors.New("no active Connection exists")
	}
	return err
}

//resolveClusterInfo resolve graphName cluster info with connection conn
func (pool *ConnectionPool) resolveClusterInfo(graphName string, conn *Connection) error {

	ctx, cancel, err := pool.NewContext(&configuration.RequestConfig{GraphName: graphName})
	defer cancel()
	if err != nil {
		return err
	}
	client := conn.GetControlClient()
	resp, err := client.GetLeader(ctx, &ultipa.GetLeaderRequest{})

	if err != nil {
		return err
	}

	if resp == nil {
		return errors.New("no resp when get leader")
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_NOT_RAFT_MODE {
		pool.IsRaft = false
		pool.GraphMgr.SetLeader(graphName, conn)
		return nil
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_RAFT_REDIRECT {
		pool.IsRaft = true
		if pool.Connections[resp.Status.ClusterInfo.Redirect] == nil {
			c, err := NewConnection(resp.Status.ClusterInfo.Redirect, pool.Config)
			if err != nil {
				return err
			}
			pool.Connections[resp.Status.ClusterInfo.Redirect] = c
		}
		pool.GraphMgr.SetLeader(graphName, pool.Connections[resp.Status.ClusterInfo.Redirect])
		pool.RefreshActives()
		return nil
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_RAFT_LEADER_NOT_YET_ELECTED {
		pool.IsRaft = true
		time.Sleep(time.Millisecond * 300)
		return utils.NewLeaderNotYetElectedError("")
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS {
		pool.IsRaft = true
		c := pool.Connections[resp.Status.ClusterInfo.LeaderAddress]
		pool.GraphMgr.SetLeader(graphName, c)
		pool.GraphMgr.ClearFollower(graphName)
		for _, follower := range resp.Status.ClusterInfo.Followers {
			fconn := pool.Connections[follower.Address]
			if fconn == nil {
				fconn2, err5 := NewConnection(follower.Address, pool.Config)
				if err5 != nil {
					continue
				}
				fconn = fconn2
				pool.Connections[follower.Address] = fconn
			}
			fconn.Host = follower.Address
			fconn.Active = follower.Status
			fconn.SetRoleFromInt32(follower.Role)
			pool.GraphMgr.AddFollower(graphName, fconn)
		}
		pool.RefreshActives()
		return nil
	} else {
		err = errors.New(fmt.Sprintf("falied to refresh cluster of graph %s: %v,%s", graphName, resp.Status.ErrorCode, resp.Status.Msg))
	}
	return err
}

// Get client by global config
func (pool *ConnectionPool) GetConn(config *configuration.UltipaConfig) (*Connection, error) {

	//if pool.Config.Consistency {
	//	return pool.GetMasterConn(config)
	//} else {
	return pool.GetRandomConn(config)
	//}
}

// Get Master of Global Graph
func (pool *ConnectionPool) GetGlobalMasterConn(config *configuration.UltipaConfig) (*Connection, error) {
	globalGraph := "global"
	if pool.GraphMgr.GetLeader(globalGraph) == nil {
		err := pool.RefreshClusterInfo(globalGraph)

		if err != nil {
			return nil, err
		}
	}

	return pool.GraphMgr.GetLeader(globalGraph), nil
}

// Get master client
func (pool *ConnectionPool) GetMasterConn(config *configuration.UltipaConfig) (*Connection, error) {

	if pool.GraphMgr.GetLeader(config.CurrentGraph) == nil {
		err := pool.RefreshClusterInfo(config.CurrentGraph)

		if err != nil {
			return nil, err
		}
	}

	return pool.GraphMgr.GetLeader(config.CurrentGraph), nil

}

//SetMasterConn (graphName , *conn) Set master client
func (pool *ConnectionPool) SetMasterConn(graphName string, conn *Connection) {
	pool.GraphMgr.SetLeader(graphName, conn)
}

// Get random client
func (pool *ConnectionPool) GetRandomConn(config *configuration.UltipaConfig) (*Connection, error) {
	pool.muActiveSafely.Lock()
	defer pool.muActiveSafely.Unlock()
	if len(pool.Actives) < 1 {
		return nil, errors.New("No Actived Connection is found")
	}

	pool.RandomTick++
	conn := pool.Actives[pool.RandomTick%len(pool.Actives)]
	if conn == nil {
		return nil, errors.New("Random Actived Connection is nil")
	}
	return conn, nil
}

// Get Task/Analytics client
func (pool *ConnectionPool) GetAnalyticsConn(config *configuration.UltipaConfig) (*Connection, error) {

	gci := pool.GraphMgr.GetGraph(config.CurrentGraph)

	if gci == nil {
		err := pool.RefreshClusterInfo(config.CurrentGraph)
		if err != nil {
			return nil, err
		}
		gci = pool.GraphMgr.GetGraph(config.CurrentGraph)
	}

	return gci.GetAnalyticConn()

}

func (pool *ConnectionPool) Close() error {
	for _, conn := range pool.Connections {
		err := conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// set context with timeout and auth info
func (pool *ConnectionPool) NewContext(config *configuration.RequestConfig) (ctx context.Context, cancel context.CancelFunc, err error) {

	if config == nil {
		config = &configuration.RequestConfig{}
	} else if config.Timezone != "" {
		_, err = time.LoadLocation(config.Timezone)
		if err != nil {
			return nil, nil, err
		}
	}

	timeout := config.Timeout

	if timeout == 0 {
		timeout = pool.Config.Timeout
	}

	if timeout == 0 {
		timeout = configuration.DefaultTimeout
	}

	if timeout < 0 {
		parentCtx := context.Background()
		ctx, cancel = context.WithCancel(parentCtx)
	} else {
		if timeout < 10 {
			timeout = 10
		}
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	}
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(pool.Config.ToContextKV(config)...))
	return ctx, cancel, nil
}

// RunHeartBeat used for special network policy for long connection(such like : force disconnection idle socket)
func (pool *ConnectionPool) RunHeartBeat() {

	if pool.Config.HeartBeat > 0 {
		go func() {
			for {
				//log.Println("Heart Beat Start... ")
				for _, conn := range pool.Connections {

					ctx, cancel, err := pool.NewContext(&configuration.RequestConfig{
						Timeout: 6,
					})
					if err != nil {
						log.Printf("heart beat failed [%s] with error :%v ", conn.Host, err)
						continue
					}
					defer cancel()
					//log.Println("Heart Beat Item", conn.Host)
					resp, err := conn.GetControlClient().SayHello(ctx, &ultipa.HelloUltipaRequest{
						Name: "go sdk refresh",
					})

					if err != nil || (resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS) {
						log.Printf("heart beat failed : ", conn.Host)
						continue
					}
				}
				//log.Println("Heart Beat End... ")
				time.Sleep(time.Duration(pool.Config.HeartBeat) * time.Second)
			}
		}()
	}
}
