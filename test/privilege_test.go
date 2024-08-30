package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func TestShowShowPrivilege(t *testing.T) {
	p, err := client.ShowPrivilege(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, privilege := range p {
		fmt.Println(privilege)
	}

	gp := &structs.GraphPrivileges{
		"default": []string{"UPDATE", "DELETE"},
		"amz":     []string{"UPDATE"},
	}

	pp := &structs.PropertyPrivileges{
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
	sp := []string{"STAT"}
	policies := []string{"yu2"}
	resp, err := client.GrantPolicy("yu", gp, sp, pp, policies, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp.Status.Code)
}
