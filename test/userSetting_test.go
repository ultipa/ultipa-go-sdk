package test

import (
	"fmt"
	"testing"
)

func TestUserSetting(t *testing.T) {
	data := "test data"
	resp, err := client.SetUserSetting("yu_tst", "s", data, nil)
	fmt.Println(resp)
	fmt.Println(err)

	resp2, err := client.GetUserSetting("yu_tst", "s", nil)
	fmt.Println(resp2)
	fmt.Println(err)
	if resp.Data != data {
		t.Errorf("GetUserSetting() got = %v, want %v", resp2.Data, data)
	}
}
