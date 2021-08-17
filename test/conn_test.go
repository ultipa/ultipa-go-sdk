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
	"ultipa-go-sdk/sdk/configuration"
)

func TestNewConn(t *testing.T) {

	//conn, err := grpc.Dial("210.13.32.146:60074", grpc.WithInsecure(), grpc.WithDefaultCallOptions())
	conn, err := grpc.Dial("210.13.32.146:60075", grpc.WithInsecure(), grpc.WithDefaultCallOptions())

	if err != nil {
		log.Fatalln(err)
	}

	client := ultipa.NewUltipaRpcsClient(conn)

	h := md5.New()
	h.Write([]byte("root"))
	pass := hex.EncodeToString(h.Sum(nil))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	ctx = metadata.AppendToOutgoingContext(ctx, "user", "root", "password", strings.ToUpper(pass), "graph_name", "multi_schema_test")

	resp, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "hello",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)

	ctx2, _ := context.WithTimeout(context.Background(), time.Second*1000)
	ctx2 = metadata.AppendToOutgoingContext(ctx2, "user", "root", "password", strings.ToUpper(pass), "graph_name", "multi_schema_test")
	resp2, err := client.Uql(ctx2, &ultipa.UqlRequest{
		Uql: "n().e().n() as path return path limit 10;",
	})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		record, err := resp2.Recv()
		if err != nil {
			log.Fatalln(err)
			break
		}
		log.Println(record.Alias, record.Paths, err)
	}

	//defer ultipa.Close()
}

func TestUql(t *testing.T) {
	client, _ := GetClient([]string{"210.13.32.146:60075"}, "default")
	res, _ := client.UQL("n().e().n() as path return path limit 10;", nil)
	log.Println(res.AliasList, res.Get(0), res.Status.Code, res.Status.Message)
}

func TestUqlWithSpecialHost(t *testing.T) {
	res, err := client.UQL("show().graph()", &configuration.RequestConfig{
		Host: "localhost:3000",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
