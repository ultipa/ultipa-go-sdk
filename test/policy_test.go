package test

import (
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/utils"
	"log"
	"testing"
)

func TestShowPolicy(t *testing.T) {
	all, err := client.ShowPolicy(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.JSONString(all))

	policy, err := client.GetPolicy("yu", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(utils.JSONString(policy))

	policy.Name = "yu_new"
	_, err = client.CreatePolicy(policy, nil)
	if err != nil {
		t.Fatal(err)
	}

	policy.Policies = []string{"yu", "yuss"}
	_, err = client.AlterPolicy(policy, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DropPolicy(policy.Name, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPolicyUql(t *testing.T) {
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

	p := structs.Policy{
		Name: "yu",
		//GraphPrivileges: gp,
		//SystemPrivileges: []string{"STAT"},
		PropertyPrivileges: pp,
		//Policies:           []string{"yu"},
	}

	log.Println(p.ToCreatePolicyUql())
	log.Println(p.ToAlterPolicyUql())

	resp, err := client.Uql(p.ToCreatePolicyUql(), nil)

	if err != nil {
		log.Println(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
	}

	resp, err = client.Uql(p.ToAlterPolicyUql(), nil)
	if err != nil {
		log.Println(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
	}

}
