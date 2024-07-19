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
	UserName           string             `json:"-"`
	PassWord           string             `json:"password,omitempty"`
	Create             string             `json:"-"`
	GraphPrivileges    GraphPrivileges    `json:"graph_privileges,omitempty"`
	SystemPrivileges   []string           `json:"system_privileges,omitempty"`
	Policies           []string           `json:"policies,omitempty"`
	PropertyPrivileges PropertyPrivileges `json:"property_privileges,omitempty"`
}

type CreateUser struct {
	UserName           string             `json:"-"`
	PassWord           string             `json:"password,omitempty"`
	GraphPrivileges    GraphPrivileges    `json:"graph_privileges,omitempty"`
	SystemPrivileges   []string           `json:"system_privileges,omitempty"`
	Policies           []string           `json:"policies,omitempty"`
	PropertyPrivileges PropertyPrivileges `json:"property_privileges,omitempty"`
}

type AlterUser struct {
	UserName           string             `json:"-"`
	PassWord           string             `json:"password,omitempty"`
	GraphPrivileges    GraphPrivileges    `json:"graph_privileges,omitempty"`
	SystemPrivileges   []string           `json:"system_privileges,omitempty"`
	Policies           []string           `json:"policies,omitempty"`
	PropertyPrivileges PropertyPrivileges `json:"property_privileges,omitempty"`
}

//func (p *User) ToCreateUserUql() string {
//	return fmt.Sprintf("create().user(\"%s\",\"%s\",\n%s,\n%s,\n%s,\n%s\n)", p.UserName, p.PassWord, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
//}
//
//func (p *User) ToAlterUserUql() string {
//	return fmt.Sprintf("alter().user(\"%s\").set({\npassword: \"%s\",\ngraph_privileges: %s,\nsystem_privileges: %s,\npolicies: %s,\nproperty_privileges: %s\n})", p.UserName, p.PassWord, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
//}

func (u *CreateUser) ToCreateUserUql() string {
	uql := fmt.Sprintf("create().user(\"%s\",\n\"%s\",\n", u.UserName, u.PassWord)

	s := ""

	if u.GraphPrivileges != nil {
		s = utils.ToJSONString(u.GraphPrivileges) + ",\n"
	} else {
		s = "{},\n"
	}
	uql += s

	if u.SystemPrivileges != nil {
		s = utils.ToJSONString(u.SystemPrivileges) + ",\n"
	} else {
		s = "[],\n"
	}
	uql += s

	if u.Policies != nil {
		s = utils.ToJSONString(u.Policies) + ",\n"
		uql += s
	} else {
		s = "[],\n"
	}
	uql += s

	if u.PropertyPrivileges != nil {
		s = utils.ToJSONString(u.PropertyPrivileges) + "\n"
	} else {
		s = "{}\n"
	}

	uql += s

	uql += ")"

	return uql
	//return fmt.Sprintf("create().user(\"%s\",\n%s,\n%s,\n%s,\n%s\n)", u.Name, utils.ToJSONString(u.GraphPrivileges), utils.ToJSONString(u.SystemPrivileges), utils.ToJSONString(u.Policies), utils.ToJSONString(u.PropertyPrivileges))
}

func (u *AlterUser) ToAlterUserUql() string {
	//return fmt.Sprintf("alter().policy(\"%s\").set({\ngraph_privileges: %s,\nsystem_privileges: %s,\npolicies: %s,\nproperty_privileges: %s\n})", p.Name, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
	return fmt.Sprintf("alter().user(\"%s\").set(%s)", u.UserName, utils.ToJSONString(u))
}
