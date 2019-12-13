package sdk

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

// Client keep the connection to ultipa db host
type Client = ultipa.UltipaRpcsClient

// ClientConn is the connection , you can close it
type ClientConn = grpc.ClientConn

// Connect a ultipa db host by hostname or ip
func Connect(username, password, host string) (_client Client, _conn *ClientConn, _err error) {
	// ultipa.SayHello()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(host, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return nil, nil, err
	}

	client := ultipa.NewUltipaRpcsClient(conn)

	return client, conn, nil

}

// TestConnect return whether connection is alive
func TestConnect(client Client) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	_, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "test",
	})

	if err != nil {
		return false, err
	}
	return true, nil
}
