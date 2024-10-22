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
	//UninstallHDCAlgo
	//removeErr := fmt.Sprintf(`remove .//algo/libs//libplugin_%s.so failed!`, algoName)
	//_, err := client.UninstallHDCAlgo(algoName, hdcName, nil)
	//if algo == nil && !strings.Contains(err.Error(), removeErr) {
	//	t.Errorf("UninstallHDCAlgo failed %v", err)
	//}
	//
	//if algo != nil && err != nil {
	//	t.Errorf("UninstallHDCAlgo failed %v", err)
	//}
	//
	//// UninstallHDCAlgo not exist again
	//_, err = client.UninstallHDCAlgo(algoName, hdcName, nil)
	//if !strings.Contains(err.Error(), removeErr) {
	//	t.Errorf("Uninstall not exist Algo failed %v", err)
	//}
	//
	//// UninstallHDCAlgo empty algoName, will success
	//_, err = client.UninstallHDCAlgo("", hdcName, nil)
	//if err != nil {
	//	t.Errorf("UninstallHDCAlgo empty algoName failed %v", err)
	//}

	// InstallHDCAlgo
	_, err := client.InstallHDCAlgos([]string{"./test_algo_lib/libplugin_lpa.so", "./test_algo_lib/lpa.yml"}, hdcName, nil)
	//files := map[string]string{
	//	"./test_algo_lib/libplugin_lpa.so": "./test_algo_lib/lpa.yml",
	//	//"./test_algo_lib/libplugin_lpa.so": "./test_algo_lib/lpa.yml",
	//}
	//_, err := client.InstallHDCAlgos(files, hdcName, nil)

	if err != nil {
		t.Errorf("InstallHDCAlgo error, %v", err)
	}

	//_, err = client.InstallHDCAlgo("./test_algo_lib/libplugin_lpa.so", "./test_algo_lib/lpa.yml", hdcName, nil)
	//versionErr := fmt.Sprintf("libplugin_%s.so:The new algo version must be greater than old!", algoName)
	//if !strings.Contains(err.Error(), versionErr) {
	//    t.Errorf("InstallHDCAlgo error, %v", err)
	//}

	//algo, _ := client.GetAlgo(algoName, nil)
	//if algo == nil {
	//    t.Error("No installed algorithm found")
	//}

	// UninstallHDCAlgo Avoid unexpected problems due to inconsistent algorithm versions and servers
	if !t.Run("UninstallHDCAlgo", TestUninstallAlgo) {
		t.Error("UninstallHDCAlgo failed")
	}
}

func TestUninstallAlgo(t *testing.T) {
	_, err := client.UninstallHDCAlgo(algoName, hdcName, nil)

	if err != nil {
		t.Fatalf("UninstallHDCAlgo error, %v", err)
	}
}

func TestRollbackHDCAlgo(t *testing.T) {
	_, err := client.RollbackHDCAlgo(algoName, hdcName, nil)

	if err != nil {
		t.Fatalf("RollbackHDCAlgo error, %v", err)
	}
}
