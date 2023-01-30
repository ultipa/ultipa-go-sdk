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

func TestCheckPropAndValueAutoData(t *testing.T) {
	var graph = "grapthInsertTest"
	hosts := []string{"192.168.1.85:61095", "192.168.1.87:61095", "192.168.1.88:61095"}
	client, _ := GetClient(hosts, graph)
	timestamp1, _ := utils.NewTimeFromString("2018-08-17T09:57:33+08:00")
	timestamp2, _ := utils.NewTimeFromString("2018-08-17 09:57:33")
	node1 := structs.Node{
		Values: &structs.Values{
			Data: map[string]interface {
			}{
				"typeTimestamp": timestamp1.GetTimeStamp(), "typeInt32": int32(1), "typeNotMatch": timestamp1.GetTimeStamp()}}, Schema: "nodeSchema2"}
	node2 := structs.Node{
		Values: &structs.Values{
			Data: map[string]interface {
			}{
				"typeTimestamp": timestamp1.GetTimeStamp(), "typeInt32": int32(1), "typeNotMatch": timestamp1.GetTimeStamp(), "typeInt32Error": int32(1)}}, Schema: "nodeSchema2"}
	node3 := structs.Node{
		Values: &structs.Values{
			Data: map[string]interface {
			}{
				"typeTimestamp": "2019-12-12 15:59:59"}}, Schema: "nodeSchema2"}
	node4 := structs.Node{
		Values: &structs.Values{
			Data: map[string]interface {
			}{}}, Schema: "nodeSchema2"}
	node5 := structs.Node{
		Values: &structs.Values{
			Data: map[string]interface {
			}{
				"typeTimestamp": "2019-12-12 15:59:59", "typeInt32": int32(1), "typeInt32Error": int32(1)}}, Schema: "nodeSchema2"}
	rows1 := []*structs.Node{&node1, &node2}
	rows2 := []*structs.Node{&node1, &node1, &node3}
	rows3 := []*structs.Node{&node1, &node1, &node4}
	rows4 := []*structs.Node{&node5}
	t.Log(timestamp2)
	cases := []struct {
		propertiesList []*structs.Property
		rows           []*structs.Node
		message        string
	}{
		{nil, rows1, "node row [1] error: values size larger than properties size."},
		{nil, rows2, "node row [2] error: values size smaller than properties size."},
		{nil, rows3, "node row [2] error: values size smaller than properties size."},
		{nil, rows4, "node row [0] error: values doesn't contain property [typeNotMatch]."},
	}
	for _, c := range cases {
		_, err1 := client.InsertNodesBatchAuto(c.rows, &configuration.InsertRequestConfig{
			InsertType: ultipa.InsertType_NORMAL})
		fmt.Println(c.rows)
		//fmt.Println(re)
		if err1.Error() != c.message {
			t.Errorf("返回信息与期望不一致，期望返回信息为%s\n实际返回信息为%s", c.message, err1.Error())
		}
	}
}