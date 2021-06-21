package connection

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

// handle all connections
type ConnectionPool struct {
	Config      *configuration.UltipaConfig
	Connections map[string]*Connection // Host : Connection
	Actives     []*Connection
	Cluster     *ClusterManager
	RandomTick int
}

func NewConnectionPool(config *configuration.UltipaConfig) *ConnectionPool {

	pool := &ConnectionPool{
		Config: config,
		Connections: map[string]*Connection{},
	}

	// Init Cluster Manager
	pool.Cluster = NewClusterManager(pool)

	// Get Connections
	pool.CreateConnections()

	// Update Raft Infos
	pool.Cluster.UltipaRaftInfo()

	// Refresh Actives
	pool.RefreshActives()

	return pool
}


func (pool *ConnectionPool) CreateConnections() error {
	var err error
	for _, host := range pool.Config.Hosts {
		conn, _ := NewConnection(host, pool.Config)
		client := conn.GetClient()
		ctx ,_ := context.WithTimeout(context.Background(), time.Duration(pool.Config.Timeout) * time.Second)
		resp, e := client.GetLeader(ctx, &ultipa.GetLeaderRequest{})

		if e != nil {
			err = e
			continue
		}

		// Not Raft Mode
		if resp.Status.ErrorCode == ultipa.ErrorCode_NOT_RAFT_MODE {
			pool.Connections[host] = conn
			pool.Cluster.Leader = conn
			return nil
		}

		//todo: raft mode
		if resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS {
			panic("Not Implement, Raft Mode Connections")
		}
	}

	return err
}

// 更新查看哪些连接还有效
func (pool *ConnectionPool) RefreshActives() {
	pool.Actives = []*Connection{}
	for _, conn := range pool.Connections {
		ctx ,_ := context.WithTimeout(context.Background(), time.Duration(pool.Config.Timeout) * time.Second)

		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(pool.Config.ToMetaKV()...))

		resp, err := conn.GetClient().SayHello(ctx, &ultipa.HelloUltipaRequest{
			Name: "go sdk refresh",
		})

		if err != nil {
			continue
		}

		if resp.Status == nil || resp.Status.ErrorCode == ultipa.ErrorCode_SUCCESS  {
			pool.Actives = append(pool.Actives, conn)
		}

	}
}

// Get client by global config
func (pool *ConnectionPool) GetConn() (*Connection, error) {
	if pool.Config.Consistency {
		return pool.GetMasterConn()
	} else {
		return pool.GetRandomConn()
	}
}

// Get master client
func (pool *ConnectionPool) GetMasterConn() (*Connection, error) {
	if pool.Cluster.Leader == nil {
		return nil, errors.New("Leader Connection is not found!")
	}

	return pool.Cluster.Leader, nil
}

// Get random client
func (pool *ConnectionPool) GetRandomConn() (*Connection, error) {
	if len(pool.Actives) < 1 {
		return nil, errors.New("No Actived Connection is found")
	}

	pool.RandomTick++

	return pool.Actives[pool.RandomTick % len(pool.Actives)], nil
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
