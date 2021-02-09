package test

import (
	"log"
	"regexp"
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

func TestUqlReplace(t *testing.T)  {
	str := `delete().nodes({"date":{"$lt":"2020-07-09"},"type": {"$gt":}2})`
	//replace_$(v) {
	//	v = v.replace(/"(\$[a-z_A-Z]+)"/g, '$1')
	//	return v
	//
	reg := regexp.MustCompile(`"(\$[a-z_A-Z]+)"`)
	println(reg.ReplaceAllString(str, "${1}"))

}