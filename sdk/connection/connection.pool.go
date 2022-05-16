package connection

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/printers"
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
	GraphMgr    *GraphManager // graph name : ClusterInfo
	Config      *configuration.UltipaConfig
	Connections map[string]*Connection // Host : Connection
	RandomTick  int
	Actives     []*Connection
	IsRaft      bool
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

func (pool *ConnectionPool) RefreshActivesWithSeconds(seconds uint32) {
	pool.Actives = []*Connection{}
	if seconds <= 0 {
		seconds = 3
	}
	for _, conn := range pool.Connections {

		ctx, _ := pool.NewContext(&configuration.RequestConfig{
			Timeout: seconds,
		})

		resp, err := conn.GetControlClient().SayHello(ctx, &ultipa.HelloUltipaRequest{
			Name: "go sdk refresh",
		})

		if err != nil {
			printers.PrintWarn(conn.Host + "failed - " + err.Error())
			conn.Active = ultipa.ServerStatus_DEAD
			continue
		}

		if resp.Status == nil || resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS {
			conn.Active = ultipa.ServerStatus_ALIVE
			pool.Actives = append(pool.Actives, conn)
		} else {
			printers.PrintWarn(conn.Host + "failed - " + resp.Status.Msg)
			conn.Active = ultipa.ServerStatus_DEAD
		}

	}
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

	var conn *Connection

	var err error
	// 如果该图集暂无初始化时
	if pool.GraphMgr.GetLeader(graphName) == nil {
		conn, err = pool.GetConn(&configuration.UltipaConfig{CurrentGraph: graphName})
	} else {
		// 已经初始化后
		conn = pool.GraphMgr.GetLeader(graphName)
	}

	if err != nil {
		return err
	}

	ctx, _ := pool.NewContext(&configuration.RequestConfig{GraphName: graphName})
	client := conn.GetControlClient()
	resp, err := client.GetLeader(ctx, &ultipa.GetLeaderRequest{})

	if resp == nil || err != nil {
		return err
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_NOT_RAFT_MODE {
		pool.IsRaft = false
		pool.GraphMgr.SetLeader(graphName, conn)
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_RAFT_REDIRECT {
		pool.IsRaft = true
		if pool.Connections[resp.Status.ClusterInfo.Redirect] == nil {
			pool.Connections[resp.Status.ClusterInfo.Redirect], err = NewConnection(resp.Status.ClusterInfo.Redirect, pool.Config)
		}

		pool.GraphMgr.SetLeader(graphName, pool.Connections[resp.Status.ClusterInfo.Redirect])
		pool.RefreshActives()
		return pool.RefreshClusterInfo(graphName)
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS {
		pool.IsRaft = true
		c := pool.Connections[resp.Status.ClusterInfo.LeaderAddress]
		pool.GraphMgr.SetLeader(graphName, c)
		pool.GraphMgr.ClearFollower(graphName)

		for _, follower := range resp.Status.ClusterInfo.Followers {
			fconn := pool.Connections[follower.Address]

			if fconn == nil {
				fconn, err = NewConnection(follower.Address, pool.Config)

				if err != nil {
					return err
				}

				pool.Connections[follower.Address] = fconn
			}

			fconn.Host = follower.Address
			fconn.Active = follower.Status
			fconn.SetRoleFromInt32(follower.Role)
			pool.GraphMgr.AddFollower(graphName, fconn)
			pool.RefreshActives()
		}
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
	if len(pool.Actives) < 1 {
		return nil, errors.New("No Actived Connection is found")
	}

	pool.RandomTick++

	return pool.Actives[pool.RandomTick%len(pool.Actives)], nil
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
func (pool *ConnectionPool) NewContext(config *configuration.RequestConfig) (context.Context, context.CancelFunc) {

	if config == nil {
		config = &configuration.RequestConfig{}
	}

	timeout := config.Timeout

	if timeout == 0 {
		timeout = pool.Config.Timeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(pool.Config.ToContextKV(config)...))
	return ctx, cancel
}

// RunHeartBeat used for special network policy for long connection(such like : force disconnection idle socket)
func (pool *ConnectionPool) RunHeartBeat() {

	if pool.Config.HeartBeat > 0 {
		go func() {
			for {
				//log.Println("Heart Beat Start... ")
				for _, conn := range pool.Connections {

					ctx, _ := pool.NewContext(&configuration.RequestConfig{
						Timeout: 6,
					})
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
