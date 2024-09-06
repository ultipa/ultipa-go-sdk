/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  index_test
 * @Date: 2022/8/4 3:41 pm
 */

package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"log"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/utils"
)

func TestCreateIndex(t *testing.T) {
	//client, _ := GetClient(hosts, graph)
	client.SetCurrentGraph("go_sdk_test")
	//resp, err := client.CreateIndex(ultipa.DBType_DBNODE, "t1", "name(10)", "name3", nil)
	//if err != nil {
	//    t.Fatal(err)
	//}

	resp, err := client.CreateNodeIndex("@t1.name(10)", "name212sssds11", nil)

	if err != nil {
		t.Error(err)
	}

	resp, err = client.CreateEdgeIndex("@e1(name(10), age)", "edgeIndex_name_age", nil)
	if err != nil {
		t.Error(err)
	}

	log.Printf(utils.JSONString(resp))
}

func TestCreateFullIndex(t *testing.T) {
	//client, _ := GetClient(hosts, graph)
	client.SetCurrentGraph("go_sdk_test")
	//resp, err := client.CreateIndex(ultipa.DBType_DBNODE, "t1", "name(10)", "name3", nil)
	//if err != nil {
	//    t.Fatal(err)
	//}

	resp, err := client.CreateNodeFullText("t1", "name", "full_name1", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err = client.CreateEdgeFullText("e1", "name", "full_name2", nil)
	if err != nil {
		t.Error(err)
	}

	log.Printf(utils.JSONString(resp))
}

func TestListIndex(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowIndex(&configuration.RequestConfig{
		GraphName: "go_sdk_test",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListNodeIndex(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowNodeIndex(nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListEdgeIndex(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowEdgeIndex(nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListFullText(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowFullText(&configuration.RequestConfig{
		GraphName: "go_sdk_test",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListNodeFullText(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowNodeFullText(nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListEdgeFullText(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowEdgeFullText(nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(utils.JSONString(indexes))
}
