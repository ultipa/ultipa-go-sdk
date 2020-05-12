package test

import (
	"testing"
	"ultipa-go-sdk/utils"
)

func TestUQL(t *testing.T) {
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
		"show().node_property()",
		"find().edges(12)",
		"find().nodes(1,2,3).select(*)", // has Nodes
		"find().edges({ _from_id : 12}).limit(3).select(*)", // has Edges
		"ab().src(12).dest(21).depth(5).limit(5).select(name)", // has Paths
		"t().n(a{age:75}).e().n({age:75}).return(a.name,a.age)", // has Attrs
		"algo().out_degree({node_id:12})", // has Values
		"show().task()", // has Tasks
	}
	for _, uql := range uqls{
		Debug("uql %v", uql)
		resUql := connet.UQL(uql)
		resJson, err := utils.StructToJSONBytes(resUql)
		if err != nil {
			t.Error(err, uqls)
		}
		Debug("\nuql res ->\n %s\n", resJson)
		if resUql.Status.Code != utils.ErrorCode_SUCCESS {
			t.Errorf("%v", resUql.Status.Code.String())
		}

	}
}