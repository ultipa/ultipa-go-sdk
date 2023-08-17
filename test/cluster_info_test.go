package test

import (
	"github.com/ultipa/ultipa-go-sdk/utils"
	"log"
	"strings"
	"testing"
)

//1.86 1.85 1.90

func TestRefreshClusterInfo(t *testing.T) {

	client, err := GetClient(hosts, graph)

	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 10; i++ {
		client.Pool.RefreshClusterInfo("global")
		utils.PrintJSON(client.Pool.GraphMgr)
	}

}

func TestSendNewGraphUQL(t *testing.T) {

}

func TestClient(t *testing.T) {

	client, err := GetClient(hosts, "global")

	if err != nil {
		log.Fatalln(err)
	}
	var connHosts []string
	for _, connection := range client.Pool.Connections {
		connHosts = append(connHosts, connection.Host)
	}
	t.Logf("connections:%s", strings.Join(connHosts, ","))

	var active []string
	for _, connection := range client.Pool.Actives {
		active = append(active, connection.Host)
	}
	t.Logf("active:%s", strings.Join(active, ","))
	client.Close()
}
