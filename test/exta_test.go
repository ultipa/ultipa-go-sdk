/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  exta_test
 * @Date: 2022/8/5 7:49 pm
 */

package test

import (
	"log"
	"testing"
)

func TestInstallExta(t *testing.T) {
	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	//client, _ := GetClient(hosts, graph)

	_, err := client.InstallExta("./test_algo_lib/libexta_page_rank.so", "./test_algo_lib/page_rank.yml", nil)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUninstallExta(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	//client, _ := GetClient(hosts, graph)

	_, err := client.UninstallExta("page_rank", nil)

	if err != nil {
		t.Error(err)
	}
}

func TestShowExta(t *testing.T) {
	extas, err := client.ShowExta(nil)
	if err == nil {
		t.Errorf("ShowExta error:%v", err)
	}
	for _, exta := range extas {
		log.Printf("%#v\n", exta)
	}
}
