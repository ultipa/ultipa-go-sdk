package main

import (
	"fmt"
	// "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk"
	// "ultipa-go-sdk/utils"
)

// import "encoding/json"

// import ultipa "ultipa-go-sdk/rpc"

func main() {
	client, conn := sdk.Connect("root", "password", "poc02.ultipa.com:60062")

	defer conn.Close()

	fmt.Println("Starting test property module")

	fmt.Printf("%v \n ", sdk.GetNodePropertyInfo(client))
	fmt.Printf("%v \n ", sdk.GetEdgePropertyInfo(client))
}
