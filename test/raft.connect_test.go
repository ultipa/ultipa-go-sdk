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
		"192.168.3.129:60161",
		"192.168.3.129:60162",
		"192.168.3.129:60163",
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
	}
}
