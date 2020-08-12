package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)
var CONCURRENT_COUNT = 100000

func TestConcurrent(t *testing.T) {
	host := "124.193.211.21:60062"
	conn, _ := GetTestDefaultConnection(&host)
	conn.DefaultConfig.GraphSetName = "default"
	conn.RefreshRaftLeader("", nil)

	var wg sync.WaitGroup
	timeStart := time.Now().UnixNano()
	for i:=1;i<CONCURRENT_COUNT;i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			uql := fmt.Sprintf("find().nodes(%d)", rand.Intn(CONCURRENT_COUNT*2))
			res := conn.UQL(uql, nil)

			if res.Status.Code == ultipa.ErrorCode_SUCCESS {
				//println(uql)
			} else {
				fmt.Printf("%s, %s", uql, utils.JSONString(res))
			}
		}()
	}
	wg.Wait()
	timeEnd := time.Now().UnixNano()
	cost := float64(float64(timeEnd - timeStart) / 1e9)
	qps := int64(float64(CONCURRENT_COUNT)/cost)
	fmt.Printf("cost: %fs, QPS: %d", cost, qps)
	fmt.Println("")

}