package sdk

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

// Connect a ultipa db host by hostname or ip
func Connect(username, password, host string) (_client Client, _conn *ClientConn, _err error) {
	// ultipa.SayHello()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(host, opts...)

	if err != nil {
		log.Printf("fail to dial: %v", err)
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
		if strings.Contains(err.Error(), "code = Unavailable") {
			return false, errors.New("Unavailable")
		}
		return false, err
	}
	return true, nil
}
