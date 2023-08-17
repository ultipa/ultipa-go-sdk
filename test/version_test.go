package test

import "testing"

func TestGetServerVersion(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	version, err := client.GetServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}
