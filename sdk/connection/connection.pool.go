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
	Graph     string
	Leader    *Connection
	Followers []*Connection
}

// handle all connections
type ConnectionPool struct {
	GraphInfos       map[string]*GraphClusterInfo // graph name : clusterinfo
	Config           *configuration.UltipaConfig
	Connections      map[string]*Connection // Host : Connection
	AnalyticsActives []*Connection
	RandomTick       int
	Actives          []*Connection
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

	return pool, nil
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

	//todo: update resp
	if resp == nil || err != nil {
		return err
	}

	if resp.Status.ErrorCode == ultipa.ErrorCode_RAFT_REDIRECT {
		pool.GraphInfos[graphName] = &GraphClusterInfo{
			Graph:  graphName,
			Leader: pool.Connections[resp.Status.ClusterInfo.Redirect],
		}

		return pool.RefreshClusterInfo(graphName)
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Msg)

	} else {

		if pool.GraphInfos[graphName] == nil {
			pool.GraphInfos[graphName] = &GraphClusterInfo{
				Graph:  graphName,
				Leader: pool.Connections[resp.Status.ClusterInfo.Redirect],
			}
		}

		for _, follower := range resp.Status.ClusterInfo.Followers {
			fconn := pool.Connections[follower.Address]
			fconn.Host = follower.Address
			fconn.Active = follower.Status
			fconn.SetRoleFromInt32(follower.Role)
			pool.GraphInfos[graphName].Followers = append(pool.GraphInfos[graphName].Followers, fconn)
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
func (pool *ConnectionPool) GetAnalyticsConn() (*Connection, error) {
	//todo
	return nil, nil
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
