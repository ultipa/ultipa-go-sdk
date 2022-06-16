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
	client, _ := GetClient([]string{"192.168.1.71:60061"}, "default")

	total := 80000000
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

		node.ID =  fmt.Sprint(total)
		node.Set("name", "abcd yefx")
		node.Set("age", int32(10))

		nodes = append(nodes, node)

		total--

		if total%30000 == 0 || total < 0 {

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
				})

				_, err := client.InsertNodesBatchBySchema(schema, nodes, &configuration.RequestConfig{
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
