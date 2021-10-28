package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/sdk/printers"
)

func TestListAlgo(t *testing.T) {

	client, _ := GetClient([]string{"192.168.1.86:40109"}, "GO_SDK_TEST")

	algos, err := client.ListAlgo(nil)

	if err != nil {
		log.Fatalln(err)
	}

	printers.PrintAlgoList(algos)
}
