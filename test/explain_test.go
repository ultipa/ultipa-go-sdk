package test

import (
	"log"
	"testing"
)

func TestExplain(t *testing.T) {

	client, _ := GetClient([]string{"192.168.1.85:54001"}, "default")


	resp, err := client.UQL("explain find().nodes() as nodes limit 1 return nodes limit 10", nil)

	if err != nil {
		log.Fatalln(err)
	}


	log.Println(resp)



}
