package connection

import (
	"ultipa-go-sdk/sdk/configuration"
)

/**
	Raft Cluster Manager, Update Raft Infos for client
 */


type ClusterManager struct {
	Leader *Connection
	Followers []*Connection
	Config *configuration.UltipaConfig
	Pool *ConnectionPool
}

func NewClusterManager(connP *ConnectionPool) *ClusterManager {
	return &ClusterManager{
		Config: connP.Config,
		Pool: connP,
	}
}
//
//func (c *ClusterManager) RefreshGraphCluster(name string) error {
//	conn, err := c.Pool.GetConn()
//	defer conn.Close()
//
//	if err != nil {
//		return err
//	}
//
//	client := conn.GetClient()
//	ctx, err := c.Pool.NewContext(nil)
//
//	if err != nil {
//		return err
//	}
//
//	resp, err := client.GetLeader(ctx, nil)
//
//	if err != nil {
//		return err
//	}
//
//	log.Fatalln(resp)
//}

//func (c *ClusterManager) Try