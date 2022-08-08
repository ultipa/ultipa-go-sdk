/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  exta_test
 * @Date: 2022/8/5 7:49 下午
 */

package test

import (
	"os"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/printers"
)

func TestInstallExta(t *testing.T) {
	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	client, _ := GetClient([]string{"192.168.1.12:60061"}, "default")

	resp, err := client.InstallExta("./test_algo_lib/libexta_page_rank.so", "./test_algo_lib/page_rank.yml", nil)

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		printers.PrintError(resp.Status.Msg)
		os.Exit(1)
	}

	if err != nil {
		printers.PrintErrAndExist(err.Error())
	}
}

func TestUninstallExta(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:60074"}, "default")
	client, _ := GetClient([]string{"192.168.1.12:60061"}, "default")

	resp, err := client.UninstallExta("page_rank", nil)

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		printers.PrintError(resp.Status.Msg)
		os.Exit(1)
	}

	if err != nil {
		printers.PrintErrAndExist(err.Error())
	}

}
