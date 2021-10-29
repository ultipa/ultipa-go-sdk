package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/sdk/printers"
)

func TestListAlgo(t *testing.T) {

	client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")

	algos, err := client.ListAlgo(nil)

	if err != nil {
		log.Fatalln(err)
	}

	printers.PrintAlgoList(algos)
}
