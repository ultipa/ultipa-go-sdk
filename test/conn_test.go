package test

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"testing"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

func TestNewConn(t *testing.T) {


	//conn, err := grpc.Dial("210.13.32.146:60074", grpc.WithInsecure(), grpc.WithDefaultCallOptions())
	conn, err := grpc.Dial("192.168.1.86:60061", grpc.WithInsecure(), grpc.WithDefaultCallOptions())

	if err != nil {
		log.Fatalln(err)
	}

	client := ultipa.NewUltipaRpcsClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 3)

	h := md5.New()
	h.Write([]byte("root"))
	pass := hex.EncodeToString(h.Sum(nil))

	ctx = metadata.AppendToOutgoingContext(ctx, "user", "root", "password", strings.ToUpper(pass), "graph_name", "default")

	resp, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "hello",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)


	//defer ultipa.Close()
}