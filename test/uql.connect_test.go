package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)


func TestUQLSingle(t *testing.T) {
	connet, _ := GetTestDefaultConnection(nil)
	uql := "find().nodes().limit(12).select(company)"
	//uql = "find().edges().limit(12).select(mark)"
	//uql = "t().n(a).e().n().limit(12).return(a.name,a.age)"
	//uql = "show().property()"
	//uql = "getUser().username(root)"
	//uql = "listGraph()"
	resUql := connet.UQL(uql)
	resJson, _ := utils.StructToJSONBytes(resUql)

	log.Printf("\nuql res ->\n %s\n", resJson)
}
func TestUQL(t *testing.T) {
	TestLogTitle("UQL")
	connet, err := GetTestDefaultConnection(nil)
	if err != nil {
		t.Error(err)
	}
	//res := connet.ListProperty(sdk.ShowPropertyRequest{Dataset: utils.DBType_DBNODE})
	//resJson, _ := utils.StructToJSONBytes(*res)
	//fmt.Printf("\nlist property -> %s\n", resJson)
	//res = connet.ListProperty(sdk.ShowPropertyRequest{Dataset: utils.DBType_DBEDGE})
	//
	//resJson, _ = utils.StructToJSONBytes(*res)
	//fmt.Printf("\nlist property -> %s\n", resJson)
	uqls := []string{
		"listGraph()",
		"listUser()", // has Tables
		"getUser().username(root)", // has Values
		"showIndex()",
		//"show().node_property()", // 待测，还不支持 05/13
		//"show().edge_property()", // 待测，还不支持 05/13
		"find().nodes(1,2,3).select(*)", // has Nodes
		"find().edges(1,2,3).select(*)", // has Edges
		//"find().nodes(1,2,3).select(*);find().edges(1,2,3).select(*)", // has Nodes amd Edges //还不支持，待测 05/13
		"ab().src(12).dest(21).depth(5).limit(5).select(*)", // has Paths // 有bug，有些node_table没有header需要 05/13
		"t().n(a).e().n(2).return(a.name,a.age)", // has Attrs // values 匹配有问题，需要服务端解决
		"t().n().e().n().select(*)",
		"showTask()",
	}
	for _, uql := range uqls {
		TestLogSubtitle("execute UQL " + uql )
		resUql := connet.UQL(uql)
		resJson, err := utils.StructToJSONBytes(resUql)
		if err != nil {
			t.Error(err, uqls)
		}
		log.Printf("\nuql res ->\n %s\n", resJson)
		if resUql.Status.Code != types.ErrorCode_SUCCESS {
			t.Errorf("%v", resUql.Status.Code.String())
		}

	}
}
