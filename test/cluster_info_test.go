package test

import (
	"testing"
)

//1.86 1.85 1.90

//func TestRefreshClusterInfo(t *testing.T) {
//
//    //client, err := GetClient(hosts, graph)
//
//    //if err != nil {
//    //	t.Fatal(err)
//    //}
//
//    for i := 0; i < 10; i++ {
//        err := client.Conn.RefreshClusterInfo("global")
//        // utils.PrintJSON(client.Conn.GraphMgr)
//        if err != nil {
//            t.Fatal(err)
//        }
//    }
//
//}

func TestSendNewGraphUQL(t *testing.T) {

}

//func TestClient(t *testing.T) {
//
//    client, err := GetClient(hosts, "global")
//
//    if err != nil {
//        t.Fatal(err)
//    }
//    var connHosts []string
//    for _, connection := range client.Conn.Connections {
//        connHosts = append(connHosts, connection.Host)
//    }
//    t.Logf("connections:%s", strings.Join(connHosts, ","))
//
//    var active []string
//    for _, connection := range client.Conn.Actives {
//        active = append(active, connection.Host)
//    }
//    t.Logf("active:%s", strings.Join(active, ","))
//    //client.Close()
//}
