package test

import (
	"log"
	"strings"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestGetLeader(t *testing.T) {
	connet, err := GetTestDefaultConnection(nil)
	if err != nil {
		t.Error(err)
	}
	res := connet.GetLeaderReuqest()
	r, _ := utils.StructToJSONString(res)
	Debug("%v", r)
}

func TestAutoRefreshLeader(t *testing.T)  {
	hosts := []string{
		"192.168.3.129:60161",
		"192.168.3.129:60162",
		"192.168.3.129:60163",
	}
	var leader string
	var followers string
	for _i, h := range hosts {
		conn, err := GetTestDefaultConnection(&h)
		if err != nil {
			t.Error(err)
		}
		err = conn.RefreshRaftLeader()
		if err != nil {
			t.Error(err)
		}
		_leader := conn.HostManager.LeaderHost
		_followers := strings.Join(conn.HostManager.FollowersHost, ",")
		log.Printf("host %v,leader %v, followers %v", h, _leader, conn.HostManager.FollowersHost)
		if _i == 0 {
			leader = _leader
			followers = _followers
		} else {
			if _leader != leader || _followers != followers {
				t.Errorf("failed host: %v", h)
			}
		}

	}
}
