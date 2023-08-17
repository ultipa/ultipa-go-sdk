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
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/configuration"
)

func TestNewConn(t *testing.T) {

	//conn, err := grpc.Dial("210.13.32.146:60074", grpc.WithInsecure(), grpc.WithDefaultCallOptions())
	conn, err := grpc.Dial(hosts[0], grpc.WithInsecure(), grpc.WithDefaultCallOptions())

	if err != nil {
		log.Fatalln(err)
	}

	client := ultipa.NewUltipaRpcsClient(conn)

	h := md5.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	ctx = metadata.AppendToOutgoingContext(ctx, "user", username, "password", strings.ToUpper(pass), graph, "multi_schema_test")

	resp, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "hello",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)

	ctx2, _ := context.WithTimeout(context.Background(), time.Second*1000)
	ctx2 = metadata.AppendToOutgoingContext(ctx2, "user", username, "password", strings.ToUpper(pass), "graph_name", "multi_schema_test")
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
	client, _ := GetClient(hosts, graph)
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

func TestRefreshPool(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	for i := 0; i < 1000; i++ {
		err := client.Pool.RefreshActivesWithSeconds(1)
		if err != nil {
			t.Log(err)
		}
		time.Sleep(time.Millisecond * 5500)
	}
}

func TestGetConnByUQL(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "show().schema()"
	_, leader, followers, global, err := client.GetConnByUQL(uql, graph)
	if err != nil {
		t.Fatal(err)
	}
	if leader == nil {
		t.Fatal("leader is nill")
	}
	if followers == nil {
		t.Fatal("followers is nill")
	}
	if global == nil {
		t.Fatal("global is nill")
	}
}

func TestConnectionSSL(t *testing.T) {

	if env["ssl_host"] == "" {
		t.Skip("no ssl host found")
		return
	}

	var err error
	config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts:        []string{env["ssl_host"]},
		Username:     env["ssl_username"],
		Password:     env["ssl_password"],
		DefaultGraph: env["ssl_graph"],
		Debug:        true,
	})

	client, err = sdk.NewUltipa(config)

	if err != nil {
		log.Fatalln(err)
	}

	uql, err := client.UQL("show().schema()", nil)
	log.Println(uql)
}
