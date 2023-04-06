package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
)

func TestBackup(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:61090"}, "default")

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.Backup("/opt/ultipa-server/exportData/backupTest/", nil)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Error("failed to backup")
	}
}
