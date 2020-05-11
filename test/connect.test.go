package main

import (
	"fmt"
	sdk "ultipa-go-sdk/sdk"
	"ultipa-go-sdk/utils"
)

func main() {
	connet := sdk.Connection{}
	host := "192.168.3.129:60162"
	//host = "localhost:60061
	err := connet.Init(host, "root", "root", "./test/ultipa.crt")
	if err != nil {
		panic(err)
	}
	defer connet.ClientInfo.Close()

	result, err1 := connet.TestConnect()
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("test %s\n", result)

	res := connet.ListProperty(sdk.ShowPropertyRequest{Dataset: sdk.DBType_DBNODE})

	resJson, _ := utils.StructToJSONBytes(*res)
	fmt.Printf("\nlist property -> %s\n", resJson)
	res = connet.ListProperty(sdk.ShowPropertyRequest{Dataset: sdk.DBType_DBEDGE})

	resJson, _ = utils.StructToJSONBytes(*res)
	fmt.Printf("\nlist property -> %s\n", resJson)

}
