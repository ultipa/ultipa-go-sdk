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

func (c *ClusterManager) UltipaRaftInfo() {
	//todo:
}

//func (c *ClusterManager) Try