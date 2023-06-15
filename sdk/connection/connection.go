package connection

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/utils"
)

type Connection struct {
	Host   string
	Conn   *grpc.ClientConn
	Client ultipa.UltipaRpcsClient
	Config *configuration.UltipaConfig
	Role   ultipa.FollowerRole // leader, follower, learner, candidate ...
	Active ultipa.ServerStatus
}

func NewConnection(host string, config *configuration.UltipaConfig) (*Connection, error) {

	var err error

	connection := &Connection{
		Config: config,
		Host:   host,
	}

	// add default mac receive size
	if config.MaxRecvSize == 0 {
		config.MaxRecvSize = 1024 * 1024 * 10
	}

	// Try to get a certificate
	certificate := utils.GetCertificate(host)
	if config.Crt == nil && certificate != nil {
		cred := credentials.NewTLS(nil)
		connection.Conn, err = grpc.Dial(host, grpc.WithTransportCredentials(cred), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.MaxRecvSize), grpc.MaxCallSendMsgSize(config.MaxRecvSize)))
	} else if config.Crt == nil {
		connection.Conn, err = grpc.Dial(host, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.MaxRecvSize), grpc.MaxCallSendMsgSize(config.MaxRecvSize)))
	} else {
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(config.Crt)
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: certPool,
		})
		connection.Conn, err = grpc.Dial(host, grpc.WithTransportCredentials(cred), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.MaxRecvSize)))
	}

	if err != nil {
		return nil, err
	}

	return connection, err
}

func (conn *Connection) GetClient() ultipa.UltipaRpcsClient {
	return ultipa.NewUltipaRpcsClient(conn.Conn)
}

func (conn *Connection) GetControlClient() ultipa.UltipaControlsClient {
	return ultipa.NewUltipaControlsClient(conn.Conn)
}

func (conn *Connection) SetRole(role ultipa.FollowerRole) {
	conn.Role = role
}

func (conn *Connection) SetRoleFromInt32(role int32) {
	conn.Role = ultipa.FollowerRole(role)
}

func (conn *Connection) HasRole(role ultipa.FollowerRole) bool {
	return (conn.Role & role) != 0
}

func (conn *Connection) Close() error {
	return conn.Conn.Close()
}
