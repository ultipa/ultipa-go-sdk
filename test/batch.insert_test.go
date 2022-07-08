package test

import (
	"fmt"
	"github.com/pieterclaerhout/go-waitgroup"
	"log"
	"math/rand"
	"testing"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func TestBatchInsertNodes(t *testing.T) {

	//client, _ := GetClient([]string{"192.168.1.85:60041"}, "zjstest")
	//client, _ := GetClient([]string{"192.168.1.71:60061"}, "default")
	client, _ := GetClient([]string{"192.168.1.85:61111"}, "batch")

	total := 10000000
	finished := 0

	wg := waitgroup.NewWaitGroup(20)

	var nodes []*structs.Node

	start := time.Now()
	rand.Seed(int64(time.Now().Second()))
	for {

		if total < 0 {
			break
		}

		node := structs.NewNode()

		t, _ := utils.NewTimeFromStringFormat("2020-10-20", "2006-01-02")
		node.ID = fmt.Sprint(total)
		node.Set("name", "abcd yefx abcd yefx")
		node.Set("age", int32(100113))
		node.Set("desc", "abcd yefx abcd yefx")
		node.Set("create_time", t.Datetime)

		nodes = append(nodes, node)

		total--

		if total%20000 == 0 || total < 0 {

			wg.BlockAdd()
			go func(nodes []*structs.Node) {
				defer wg.Done()
				schema := structs.NewSchema("default")
				schema.Properties = append(schema.Properties, &structs.Property{
					Name: "name",
					Type: ultipa.PropertyType_STRING,
				}, &structs.Property{
					Name: "age",
					Type: ultipa.PropertyType_INT32,
				}, &structs.Property{
					Name: "desc",
					Type: ultipa.PropertyType_STRING,
				},
				&structs.Property{
					Name: "create_time",
					Type: ultipa.PropertyType_DATETIME,
				},
				)

				_, err := client.InsertNodesBatchBySchema(schema, nodes, &configuration.InsertRequestConfig{
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

}

func TestBatchInsertEdges(t *testing.T) {

	//client, _ := GetClient([]string{"192.168.1.85:60041"}, "zjstest")
	//client, _ := GetClient([]string{"192.168.1.71:60061"}, "default")
	client, _ := GetClient([]string{"192.168.1.85:61111"}, "batch")

	total := 10000000
	finished := 0

	wg := waitgroup.NewWaitGroup(20)

	var edges []*structs.Edge

	schema := structs.NewSchema("default")
	schema.Properties = append(schema.Properties, &structs.Property{
		Name: "create_time",
		Type: ultipa.PropertyType_DATETIME,
	}, &structs.Property{
		Name: "type",
		Type: ultipa.PropertyType_STRING,
	}, &structs.Property{
		Name: "invest",
		Type: ultipa.PropertyType_INT32,
	})

	start := time.Now()
	rand.Seed(int64(time.Now().Second()))
	for {

		if total < 0 {
			break
		}

		edge := structs.NewEdge()
		edge.From = fmt.Sprint(total)
		edge.To = fmt.Sprint(total - 1)

		edge.Set("type", "tyoe12313kjhasdkahd")
		edge.Set("create_time", utils.TimeToUint64(time.Now()))
		edge.Set("invest", int32(12321))


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
