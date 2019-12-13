package main

import (
	"fmt"
	"time"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/algorithm"
	// "ultipa-go-sdk/utils"
)

func main() {
	fmt.Println("test louvain sdk")

	client, conn := sdk.Connect("root", "password", "192.168.3.17:60061")

	defer conn.Close()

	params := algorithm.NewLouvainParams()
	msg := algorithm.StartLouvainTask(client, params)

	fmt.Printf("%v \n", msg)

	for {
		time.Sleep(time.Second * 2)

		msg := algorithm.GetTask(client)
		fmt.Printf("%v \n", msg)

	}

}
