package test

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"testing"
)

func TestUser(t *testing.T) {
	users, err := client.ShowUser(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}

	user, err := client.GetUser("yu", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)

}

func TestUserUql(t *testing.T) {
	//gp := structs.GraphPrivileges{
	//	"default": []string{"UPDATE", "DELETE"},
	//	"amz":     []string{"UPDATE"},
	//}
	pp := structs.PropertyPrivileges{
		"node": {
			"read":  {},
			"write": {{"default", "*", "*"}, {"amz", "nodx", "age"}},
			"deny":  {},
		},
		"edge": {
			"read":  {},
			"write": {{"default", "*", "*"}},
			"deny":  {},
		},
	}

	p := structs.User{
		UserName: "yu1112",
		PassWord: "1112",
		//GraphPrivileges: gp,
		//SystemPrivileges: []string{"STAT"},
		PropertyPrivileges: pp,
		//Policies:           []string{"yu"},
	}

	fmt.Println(p.ToCreateUserUql())
	fmt.Println(p.ToAlterUserUql())
}
