package sdk

import (
	"google.golang.org/grpc"
	"log"
	ultipa "ultipa-go-sdk/rpc"
)

func Connect(username, password, host string) (ultipa.UltipaRpcsClient, *grpc.ClientConn) {
	// ultipa.SayHello()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(host, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := ultipa.NewUltipaRpcsClient(conn)

	return client, conn

}
