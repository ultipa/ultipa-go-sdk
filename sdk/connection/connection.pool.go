package connection

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

type GraphClusterInfo struct {
	Graph         string
	Leader        *Connection
	Followers     []*Connection
	Algos         []*Connection
	LastAlgoIndex int //记录上次使用的 Task 节点索引
}

func (gci *GraphClusterInfo) AddFollower(conn *Connection) {

	gci.Graph = conn.Host

	if gci.HasConn(conn) == true {
		return
	}

	gci.Followers = append(gci.Followers, conn)

	if conn.HasRole(ultipa.FollowerRole_ROLE_ALGO_EXECUTABLE) {
		gci.Algos = append(gci.Algos, conn)
	}

}

func (gci *GraphClusterInfo) ClearFollower() {
	gci.Followers = []*Connection{}
	gci.Algos = []*Connection{}
}

func (gci *GraphClusterInfo) GetAnalyticConn() (*Connection, error) {

	if len(gci.Algos) == 0 {
		return nil, errors.New("no Algo/Task Instance Found")
	}

	gci.LastAlgoIndex++

	return gci.Algos[gci.LastAlgoIndex%len(gci.Algos)], nil
}

func (gci *GraphClusterInfo) HasConn(_conn *Connection) bool {

	if gci.Leader == _conn {
		return true
	}

	for _, conn := range gci.Followers {
		if conn == _conn {
			return true
		}
	}

	return false
}

// handle all connections
type ConnectionPool struct {
	GraphInfos  map[string]*GraphClusterInfo // graph name : ClusterInfo
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
		GraphInfos:  map[string]*GraphClusterInfo{},
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

// 更新查看哪些连接还有效
func (pool *ConnectionPool) RefreshActives() {
	pool.Actives = []*Connection{}
	for _, conn := range pool.Connections {

		ctx, _ := pool.NewContext(nil)

		resp, err := conn.GetClient().SayHello(ctx, &ultipa.HelloUltipaRequest{
			Name: "go sdk refresh",
		})

		if err != nil {
			continue
		}

		if resp.Status == nil || resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS {
			pool.Actives = append(pool.Actives, conn)
		}

	}
}

// sync cluster info from server
func (pool *ConnectionPool) RefreshClusterInfo(graphName string) error {

	var conn *Connection
	var err error
	if pool.GraphInfos[graphName] == nil || pool.GraphInfos[graphName].Leader == nil {
		conn, err = pool.GetConn(nil)
	} else {
		conn = pool.GraphInfos[graphName].Leader
	}

	if err != nil {
		return err
	}

	ctx, _ := pool.NewContext(&configuration.RequestConfig{GraphName: graphName})
	client := conn.GetClient()
	resp, err := client.GetLeader(ctx, &ultipa.GetLeaderRequest{})

	if resp == nil || err != nil {
		return err
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_NOT_RAFT_MODE {
		pool.IsRaft = false
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_RAFT_REDIRECT {
		pool.IsRaft = true
		if pool.Connections[resp.Status.ClusterInfo.Redirect] == nil {
			pool.Connections[resp.Status.ClusterInfo.Redirect], err = NewConnection(resp.Status.ClusterInfo.Redirect, pool.Config)
		}

		pool.GraphInfos[graphName] = &GraphClusterInfo{
			Graph:  graphName,
			Leader: pool.Connections[resp.Status.ClusterInfo.Redirect],
		}

		return pool.RefreshClusterInfo(graphName)
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Msg)
	} else {
		pool.IsRaft = true
		if pool.GraphInfos[graphName] == nil {
			pool.GraphInfos[graphName] = &GraphClusterInfo{
				Graph:  graphName,
				Leader: pool.Connections[resp.Status.ClusterInfo.Redirect],
			}
		}

		pool.GraphInfos[graphName].ClearFollower()

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
			pool.GraphInfos[graphName].AddFollower(fconn)
		}
	}

	return err
}

// Get client by global config
func (pool *ConnectionPool) GetConn(config *configuration.UltipaConfig) (*Connection, error) {

	if pool.Config.Consistency {
		return pool.GetMasterConn(config)
	} else {
		return pool.GetRandomConn(config)
	}
}

// Get master client
func (pool *ConnectionPool) GetMasterConn(config *configuration.UltipaConfig) (*Connection, error) {

	if pool.GraphInfos[config.CurrentGraph] == nil || pool.GraphInfos[config.CurrentGraph].Leader == nil {
		err := pool.RefreshClusterInfo(config.CurrentGraph)

		if err != nil {
			return nil, err
		}
	}

	return pool.GraphInfos[config.CurrentGraph].Leader, nil

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

	gci := pool.GraphInfos[config.CurrentGraph]

	if gci == nil {
		err := pool.RefreshClusterInfo(config.CurrentGraph)
		if err != nil {
			return nil, err
		}
		gci = pool.GraphInfos[config.CurrentGraph]
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(pool.Config.Timeout)*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(pool.Config.ToContextKV(config)...))
	return ctx, cancel
}
