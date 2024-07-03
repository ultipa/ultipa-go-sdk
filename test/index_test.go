/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  index_test
 * @Date: 2022/8/4 3:41 下午
 */

package test

import (
	"github.com/ultipa/ultipa-go-sdk/utils"
	"log"
	"testing"
)

func TestListIndex(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowIndex(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListNodeIndex(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowNodeIndex(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListEdgeIndex(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowEdgeIndex(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListFullText(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowFullText(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListNodeFullText(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowNodeFullText(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListEdgeFullText(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	indexes, err := client.ShowEdgeFullText(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}
