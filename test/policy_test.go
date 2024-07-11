package test

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/utils"
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

	p := policy.Policies[0]
	p.Name = "yu_new"
	_, err = client.CreatePolicy(p, nil)
	if err != nil {
		t.Fatal(err)
	}

	p.Policies = []string{"yu", "yuss"}
	_, err = client.AlterPolicy(p, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DropPolicy(p.Name, nil)
	if err != nil {
		t.Fatal(err)
	}
}
