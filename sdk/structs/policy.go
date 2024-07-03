package structs

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
	Name               string
	GraphPrivileges    graphPrivileges
	SystemPrivileges   []string
	PropertyPrivileges propertyPrivileges
	Policies           []string
}

type graphPrivileges map[string][]string

type propertyPrivileges map[string]map[string][][]string
