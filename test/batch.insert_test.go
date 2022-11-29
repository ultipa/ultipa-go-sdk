package test

import (
	"fmt"
	"github.com/pieterclaerhout/go-waitgroup"
	"log"
	"math/rand"
	"testing"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/api"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
)

func TestBatchInsertNodes(t *testing.T) {

	//client, _ := GetClient([]string{"192.168.1.85:60041"}, "zjstest")
	//client, _ := GetClient([]string{"192.168.1.71:60061"}, "default")
	conn, _ := GetClient([]string{"192.168.1.85:64801", "192.168.1.85:64802", "192.168.1.85:64803"}, "default")
	schema := "text_schema"
	createSchema(t, schema, conn)
	batchInsert(schema, conn)
	checkInsertionResult(t, conn, schema)
}

func batchInsert(schema string, conn *api.UltipaAPI) []*structs.Node {
	total := 500
	finished := 0

	wg := waitgroup.NewWaitGroup(20)

	var nodes []*structs.Node

	start := time.Now()
	rand.Seed(int64(time.Now().Second()))
	for {

		if total < 1 {
			break
		}

		node := structs.NewNode()

		node.ID = fmt.Sprint(total)
		value := rand.Intn(1000)
		node.Set("username", fmt.Sprintf("用户_%d", value))
		node.Set("password", fmt.Sprintf("password_%d", value))

		nodes = append(nodes, node)

		total--

		if total%30000 == 0 || total < 0 {

			wg.BlockAdd()
			go func(nodes []*structs.Node) {
				defer wg.Done()
				schema := structs.NewSchema(schema)
				schema.Properties = append(schema.Properties, &structs.Property{
					Name: "username",
					Type: ultipa.PropertyType_STRING,
				}, &structs.Property{
					Name: "password",
					Type: ultipa.PropertyType_TEXT,
				})

				_, err := conn.InsertNodesBatchBySchema(schema, nodes, &configuration.InsertRequestConfig{
					InsertType: ultipa.InsertType_OVERWRITE,
				})

				finished += len(nodes)
				if err != nil {
					log.Println(err)
					return
				}

				log.Printf("finished: %v, speed: %v", finished, float64(finished)/time.Since(start).Seconds())
			}(nodes)

			nodes = []*structs.Node{}
		}
	}

	wg.Wait()
	return nodes
}

func createSchema(t *testing.T, schema string, conn *api.UltipaAPI) {
	newSchemaWithProperties := &structs.Schema{
		Name: schema,
		Desc: "A Schema with 2 properties",
		Properties: []*structs.Property{
			{
				Name: "username",
				Type: ultipa.PropertyType_STRING,
			},
			{
				Name: "password",
				Type: ultipa.PropertyType_TEXT,
			},
		},
	}

	resp2, err := conn.CreateSchema(newSchemaWithProperties, true, nil)
	if err != nil {
		t.Error("failed to create schema", err)
	}
	log.Println(resp2)
}

func checkInsertionResult(t *testing.T, conn *api.UltipaAPI, schema string) {
	resp3, err := conn.UQL(fmt.Sprintf("find().nodes({@%s}) as nodes return nodes{*}", schema), nil)
	if err != nil {
		t.Errorf("failed to query insertion result. %v", err)
	}
	log.Println(resp3)
	nodes, schemas, err := resp3.Alias("nodes").AsNodes()
	printers.PrintNodes(nodes, schemas)
	if len(nodes) != 500 {
		t.Errorf("expected 500, got %d", len(nodes))
	}
}

func TestBatchInsertEdges(t *testing.T) {

	//client, _ := GetClient([]string{"192.168.1.85:60041"}, "zjstest")
	//client, _ := GetClient([]string{"192.168.1.71:60061"}, "default")
	client, _ := GetClient([]string{"192.168.1.85:61115"}, "gongshang")

	total := 500
	finished := 0

	wg := waitgroup.NewWaitGroup(10)

	var edges []*structs.Edge

	schema := structs.NewSchema("default")
	schema.Properties = append(schema.Properties, &structs.Property{
		Name: "e1",
		Type: ultipa.PropertyType_STRING,
	}, &structs.Property{
		Name: "e2",
		Type: ultipa.PropertyType_STRING,
	}, &structs.Property{
		Name: "e3",
		Type: ultipa.PropertyType_STRING,
	}, &structs.Property{
		Name: "e4",
		Type: ultipa.PropertyType_STRING,
	}, &structs.Property{
		Name: "e5",
		Type: ultipa.PropertyType_STRING,
	})

	start := time.Now()
	rand.Seed(int64(time.Now().Second()))
	for {

		if total < 1 {
			break
		}

		edge := structs.NewEdge()
		edge.From = fmt.Sprint(total)
		edge.To = fmt.Sprint(total - 1)

		edge.Set("e1", "abc")
		edge.Set("e2", "abc")
		edge.Set("e3", "abc")
		edge.Set("e4", "abc")
		edge.Set("e5", "abc")

		edges = append(edges, edge)

		total--

		if total%30000 == 0 || total < 0 {

			wg.BlockAdd()
			go func(edges []*structs.Edge) {
				defer wg.Done()

				_, err := client.InsertEdgesBatchBySchema(schema, edges, &configuration.InsertRequestConfig{
					InsertType: ultipa.InsertType_OVERWRITE,
				})

				finished += len(edges)
				if err != nil {
					log.Println(err)
					return
				}

				log.Printf("finished: %v, speed: %v", finished, float64(finished)/time.Since(start).Seconds())
			}(edges)

			edges = []*structs.Edge{}
		}
	}

	wg.Wait()

}
