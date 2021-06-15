package connection

import (
	"google.golang.org/grpc"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

type Connection struct {
	Host string
	Conn *grpc.ClientConn
	Client ultipa.UltipaRpcsClient
	Config *configuration.UltipaConfig
	Active bool
}

func NewConnection(host string, config *configuration.UltipaConfig) (*Connection, error) {

	var err error

	connection :=  &Connection{
		Config: config,
	}

	// add default mac receive size
	if config.MaxRecvSize == 0 { config.MaxRecvSize = 1024 * 1024 * 10 }

	if config.Crt == nil {
		connection.Conn, err = grpc.Dial(host, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.MaxRecvSize)))
	} else {
		connection.Conn, err = grpc.Dial(host, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.MaxRecvSize)))
	}

	if err != nil { return nil, err}

	return connection, err
}

func (conn *Connection) GetClient() ultipa.UltipaRpcsClient {
	return ultipa.NewUltipaRpcsClient(conn.Conn)
}

func (conn *Connection) Close() error {
	return conn.Conn.Close()
}