package test

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"testing"
)

func TestUserSetting(t *testing.T) {
	key := "yu_tst"
	data := "test data"
	resp, err := client.SetUserSetting(&structs.SetUserSetting{
		UserName: key,
		Type:     "int",
		Data:     data,
	}, nil)
	fmt.Println(resp)
	fmt.Println(err)

	resp2, err := client.GetUserSetting(&structs.GetUserSetting{
		UserName: key,
		//Type:     "int",
	}, nil)
	fmt.Println(err)

	fmt.Println(resp2)
}
