/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  index_test
 * @Date: 2022/8/4 3:41 下午
 */

package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestListIndex(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	indexes, err := client.ListIndex(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListNodeIndex(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	indexes, err := client.ListNodeIndex(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListEdgeIndex(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	indexes, err := client.ListEdgeIndex(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListFullText(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	indexes, err := client.ListFullText(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListNodeFullText(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	indexes, err := client.ListNodeFullText(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}

func TestListEdgeFullText(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	indexes, err := client.ListEdgeFullText(nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(utils.JSONString(indexes))
}
