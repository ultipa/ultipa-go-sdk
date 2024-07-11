package structs

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

// Policy create().policy(
//
//	  "manager",
//	 {"*": ["UPDATE", "DELETE"]},
//	 ["STAT"],
//	 ["yu", "yu2"],
//	  {
//	  "node": {
//	    "write": [["default","*","*"], ["amz","nodx","age"]]
//	  },
//	  "edge": {
//	    "write": [["default","*","*"]]
//	  }
//	}
//
// )
type Policy struct {
	Name               string             `json:"-"`
	GraphPrivileges    graphPrivileges    `json:"graph_privileges,omitempty" json:"graph_privileges,omitempty"`
	SystemPrivileges   []string           `json:"system_privileges,omitempty"`
	PropertyPrivileges propertyPrivileges `json:"property_privileges,omitempty"`
	Policies           []string           `json:"policies,omitempty"`
}

/*
	{
	  "node": {
	    "read": [],
	    "write": [
	      ["default", "*", "*"],
	      ["amz", "nodx", "age"]
	    ],
	    "deny": []
	  },
	  "edge": {
	    "read": [],
	    "write": [
	      ["default", "*", "*"]
	    ],
	    "deny": []
	  }
	}
*/
type propertyPrivileges map[string]map[string][][]string

type graphPrivileges map[string][]string

func (p *Policy) ToCreatePolicyUql() string {
	return fmt.Sprintf("create().policy(\"%s\",\n%s,\n%s,\n%s,\n%s\n)", p.Name, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
}

func (p *Policy) ToAlterPolicyUql() string {
	return fmt.Sprintf("alter().policy(\"%s\").set({\ngraph_privileges: %s,\nsystem_privileges: %s,\npolicies: %s,\nproperty_privileges: %s\n})", p.Name, utils.ToJSONString(p.GraphPrivileges), utils.ToJSONString(p.SystemPrivileges), utils.ToJSONString(p.Policies), utils.ToJSONString(p.PropertyPrivileges))
	//return fmt.Sprintf("alter().policy(\"%s\").set({\n%s\n})", p.Name, utils.ToJSONString(p))
}
