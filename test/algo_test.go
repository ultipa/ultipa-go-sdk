package test

import (
	"log"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
)

func TestShowAlgo(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	algos, err := client.ShowAlgo(nil)

	if err != nil {
		log.Fatalln(err)
	}

	printers.PrintAlgoList(algos)
}

func TestAlgo(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	algoName := "lpa"
	_, err := client.GetAlgo(algoName, nil)

	if err == nil {
		// if algo exist UninstallAlgo
		_, err := client.UninstallAlgo("lpa", nil)
		if err != nil {
			t.Errorf("UninstallAlgo error, %v", err)
		}
	}

	// InstallAlgo
	_, err = client.InstallAlgo("./data/installAlgo/libplugin_lpa.so", "./data/installAlgo/lpa.yml", nil)

	if err != nil {
		t.Errorf("InstallAlgo error, %v", err)
	}

	algo, _ := client.GetAlgo(algoName, nil)
	if algo == nil {
		t.Errorf("No installed algorithm found")
	}

	// UninstallAlgo
	_, err = client.UninstallAlgo("lpa", nil)
	if err != nil {
		t.Errorf("UninstallAlgo error, %v", err)
	}
}
