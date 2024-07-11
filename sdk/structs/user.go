package structs

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

// PrivilegeToUser
// create().user(
// "user01",
// "pwABC123",
// {"*": ["UPDATE","ALGO","LTE","UFE"]},
// ["STAT","TOP","KILL"],
// [],
// {
// "node": {
// "write": [["*","*","*"]]
// },
// "edge": {
// "write": [["*","*","*"]]
// }
// }
// )
type User struct {
	UserName           string
	PassWord           string
	GraphPrivileges    graphPrivileges
	SystemPrivileges   []string
	Policies           []string
	PropertyPrivileges propertyPrivileges
}

func (p *User) ToCreateUserUql() string {
	return fmt.Sprintf("create().user(\"%s\",\"%s\",\n%s,\n%s,\n%s,\n%s\n)", p.UserName, p.PassWord, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
}

func (p *User) ToAlterUserUql() string {
	return fmt.Sprintf("alter().user(\"%s\").set({\npassword: \"%s\",\ngraph_privileges: %s,\nsystem_privileges: %s,\npolicies: %s,\nproperty_privileges: %s\n})", p.UserName, p.UserName, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
}
