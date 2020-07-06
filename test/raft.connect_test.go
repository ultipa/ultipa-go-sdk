package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestGetLeader(t *testing.T) {

	TestLogTitle("Get Leader")
	connet, err := GetTestDefaultConnection(nil)
	if err != nil {
		t.Error(err)
	}
	res := connet.GetLeaderReuqest(nil)
	r, _ := utils.StructToJSONString(res)
	log.Printf("%v", r)

	err = connet.RefreshRaftLeader("", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestAutoRefreshLeader(t *testing.T) {
	hosts := []string{
		"124.193.211.21:60161",
		"124.193.211.21:60162",
		"124.193.211.21:60163",
	}
	for _, h := range hosts {
		conn, err := GetTestDefaultConnection(&h)
		if err != nil {
			t.Error(err)
		}
		err = conn.RefreshRaftLeader("", nil)
		if err != nil {
			t.Error(err)
		}
		res := conn.GetLeaderReuqest(nil)
		r, _ := utils.StructToJSONString(res)
		log.Printf("%v", r)
	}
}
