package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestStat(t *testing.T) {
	conn, _ := GetTestDefaultConnection(nil)

	res := conn.Stat(nil)
	resJson, _ := utils.StructToPrettyJSONString(res)

	log.Printf("\nuql res ->\n %s\n", resJson)
}
func TestCluserInfo(t *testing.T)  {
	conn, _ := GetTestDefaultConnection(nil)

	res := conn.ClusterInfo(nil)
	resJson, _ := utils.StructToPrettyJSONString(res)

	log.Printf("\nuql res ->\n %s\n", resJson)
}