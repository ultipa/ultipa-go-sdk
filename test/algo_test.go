package test

import (
	"testing"
)

func TestShowAlgo(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

	algos, err := client.ShowAlgo(nil)

	if err != nil {
		t.Fatal(err)
	}
	if len(algos) == 0 {
		t.Log("no algo return")
	}
	//printers.PrintAlgoList(algos)
}

var (
	algoName = "lpa"
	hdcName  = "hdc-server-1"
)

func TestAlgo(t *testing.T) {
	//algo, _ := client.GetAlgo(algoName, nil)
	//
	//UninstallAlgo
	//removeErr := fmt.Sprintf(`remove .//algo/libs//libplugin_%s.so failed!`, algoName)
	//_, err := client.UninstallAlgo(algoName, hdcName, nil)
	//if algo == nil && !strings.Contains(err.Error(), removeErr) {
	//	t.Errorf("UninstallAlgo failed %v", err)
	//}
	//
	//if algo != nil && err != nil {
	//	t.Errorf("UninstallAlgo failed %v", err)
	//}
	//
	//// UninstallAlgo not exist again
	//_, err = client.UninstallAlgo(algoName, hdcName, nil)
	//if !strings.Contains(err.Error(), removeErr) {
	//	t.Errorf("Uninstall not exist Algo failed %v", err)
	//}
	//
	//// UninstallAlgo empty algoName, will success
	//_, err = client.UninstallAlgo("", hdcName, nil)
	//if err != nil {
	//	t.Errorf("UninstallAlgo empty algoName failed %v", err)
	//}

	// InstallAlgo
	_, err := client.InstallAlgo("./test_algo_lib/libplugin_lpa.so", "./test_algo_lib/lpa.yml", hdcName, nil)

	if err != nil {
		t.Errorf("InstallAlgo error, %v", err)
	}

	//_, err = client.InstallAlgo("./test_algo_lib/libplugin_lpa.so", "./test_algo_lib/lpa.yml", hdcName, nil)
	//versionErr := fmt.Sprintf("libplugin_%s.so:The new algo version must be greater than old!", algoName)
	//if !strings.Contains(err.Error(), versionErr) {
	//    t.Errorf("InstallAlgo error, %v", err)
	//}

	//algo, _ := client.GetAlgo(algoName, nil)
	//if algo == nil {
	//    t.Error("No installed algorithm found")
	//}

	// UninstallAlgo Avoid unexpected problems due to inconsistent algorithm versions and servers
	if !t.Run("UninstallAlgo", TestUninstallAlgo) {
		t.Error("UninstallAlgo failed")
	}
}

func TestUninstallAlgo(t *testing.T) {
	_, err := client.UninstallAlgo(algoName, hdcName, nil)

	if err != nil {
		t.Fatalf("UninstallAlgo error, %v", err)
	}
}

func TestRollbackHDCAlgo(t *testing.T) {
	_, err := client.RollbackHDCAlgo(algoName, hdcName, nil)

	if err != nil {
		t.Fatalf("RollbackHDCAlgo error, %v", err)
	}
}
