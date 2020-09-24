package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

func TestUser(t *testing.T) {
	t.Skip("skip")
	TestLogTitle("UQL User Policy")
	connet, err := GetTestDefaultConnection(nil)

	if err != nil {
		t.Error(err)
	}
	uqls := []string{
		`listPolicy()`,
		`listPrivilege()`,
		`listUser()`,
		`deleteUser().username("autotest_username")`,
		`deletePolicy().name("autotest_policy")`,
		`listPolicy()`,
		`listPrivilege()`,
		`listUser()`,
		`createPolicy().name("autotest_policy").graph_privileges([{default:["UFE","LTE","DROP_PROPERTY"]}]).system_privileges(["GRAPH","POLICY","USER"]).policies(["DB"])`,
		`getPolicy().name("autotest_policy")`,
		`updatePolicy().name("autotest_policy").graph_privileges([{default:[]}]).system_privileges([]).policies([])`,
		`getPolicy().name("autotest_policy")`,
		`updatePolicy().name("autotest_policy").graph_privileges([{default:["ALGO","UQL"]}]).system_privileges(["POLICY"]).policies(["ROOT"])`,
		`getPolicy().name("autotest_policy")`,
		`createUser().username("autotest_username").password("autotest_username").graph_privileges([{default:["LTE"]}]).system_privileges(["USER"]).policies(["autotest_policy"])`,
		`getUser().username("autotest_username")`,
		`updateUser().username("autotest_username").graph_privileges([{default:[""]}]).system_privileges([""]).policies([""])`,
		`getUser().username("autotest_username")`,
		`grant().username("autotest_username").graph_privileges([{*:["UFE","LTE","DROP_PROPERTY"]}]).system_privileges(["GRAPH","POLICY","USER"]).policies(["ROOT"])`,
		`getUser().username("autotest_username")`,
		`revoke().username("autotest_username").graph_privileges([{*:["UFE","LTE","DROP_PROPERTY"]}]).system_privileges(["GRAPH","POLICY","USER"]).policies(["ROOT"])`,
		`getUser().username("autotest_username")`,
	}
	for _, uql := range uqls {
		TestLogSubtitle("execute UQL " + uql)
		resUql := connet.UQL(uql, nil)
		resJson, err := utils.StructToPrettyJSONString(resUql)
		if err != nil {
			t.Error(err, uqls)
		}
		log.Printf("\nuql res ->\n %s\n", resJson)
		if resUql.Status.Code != types.ErrorCode_SUCCESS {
			t.Errorf("%v", resUql.Status.Code.String())
		}

	}
}
