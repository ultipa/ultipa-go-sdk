package test

import (
	"log"
	"os"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/printers"
)

func TestListAlgo(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	client, _ := GetClient([]string{"192.168.1.60:60061"}, "default")

	algos, err := client.ShowAlgo(nil)

	if err != nil {
		log.Fatalln(err)
	}

	printers.PrintAlgoList(algos)
}

func TestInstallAlgo(t *testing.T) {
	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	client, _ := GetClient([]string{"192.168.1.60:60061"}, "default")

	resp, err := client.InstallAlgo("./test_algo_lib/libplugin_lpa.so", "./test_algo_lib/lpa.yml", nil)

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		printers.PrintError(resp.Status.Msg)
		os.Exit(1)
	}

	if err != nil {
		printers.PrintErrAndExist(err.Error())
	}

	TestListAlgo(t)
}

func TestUninstallAlgo(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	client, _ := GetClient([]string{"192.168.1.60:60061"}, "default")

	resp, err := client.UninstallAlgo("lpa", nil)

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		printers.PrintError(resp.Status.Msg)
		os.Exit(1)
	}

	if err != nil {
		printers.PrintErrAndExist(err.Error())
	}

	TestListAlgo(t)

}
