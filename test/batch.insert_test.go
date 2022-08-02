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
)

func TestBatchInsertNodes(t *testing.T) {

	//client, _ := GetClient([]string{"192.168.1.85:60041"}, "zjstest")
	//client, _ := GetClient([]string{"192.168.1.71:60061"}, "default")
	client, _ := GetClient([]string{"192.168.1.85:61115"}, "gongshang")

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
		node.Set("n2", "abcd yefx abcd yefx")
		node.Set("n1", int32(10))

		nodes = append(nodes, node)

		total--

		if total%30000 == 0 || total < 0 {

			wg.BlockAdd()
			go func(nodes []*structs.Node) {
				defer wg.Done()
				schema := structs.NewSchema("default")
				schema.Properties = append(schema.Properties, &structs.Property{
					Name: "n2",
					Type: ultipa.PropertyType_STRING,
				}, &structs.Property{
					Name: "n1",
					Type: ultipa.PropertyType_INT32,
				})

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
