package test

import "testing"

func TestGetServerVersion(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.87:61095"}, "miniCircle")
	version, err := client.GetServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}
