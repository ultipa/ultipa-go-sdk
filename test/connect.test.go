package main

import (
	"fmt"
	"log"
	sdk "ultipa-go-sdk/sdk"
	"ultipa-go-sdk/utils"
)

func main() {
	connet := sdk.Connection{}
	host := "192.168.3.129:60162"
	host = "localhost:60061"
	err := connet.Init(host, "root", "root", "./test/ultipa.crt")
	if err != nil {
		panic(err)
	}

	result, err1 := connet.TestConnect()
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("test %s\n", result)

	res := connet.ListProperty(sdk.ShowPropertyRequest{Dataset: utils.DBType_DBNODE})

	resJson, _ := utils.StructToJSONBytes(*res)
	fmt.Printf("\nlist property -> %s\n", resJson)
	res = connet.ListProperty(sdk.ShowPropertyRequest{Dataset: utils.DBType_DBEDGE})

	resJson, _ = utils.StructToJSONBytes(*res)
	fmt.Printf("\nlist property -> %s\n", resJson)

	uqls := []string{
		"show().node_property()",
		"find().edges(12)",
		"find().nodes(1,2,3).select(*)", // has Nodes
		"find().edges({ _from_id : 12}).limit(3).select(*)", // has Edges
		"ab().src(12).dest(21).depth(5).limit(5).select(name)", // has Paths
		"t().n(a{age:75}).e().n({age:75}).return(a.name,a.age)", // has Attrs
		"algo().out_degree({node_id:12})",
	}
	for _, uql := range uqls{
		log.Printf(uql)
		resUql := connet.UQL(uql)
		resJson, _ = utils.StructToJSONBytes(resUql)
		fmt.Printf("\nuql test -> %s\n", resJson)
	}

}
