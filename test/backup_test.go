package test

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"log"
	"testing"
)

func TestBackup(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.Backup("/opt/ultipa-server/exportData/backupTest/", nil)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Error("failed to backup")
	}
}
