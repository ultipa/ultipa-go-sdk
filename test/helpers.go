package test

import (
	"log"
	"math/rand"
	"time"
	"ultipa-go-sdk/sdk"
)
var hosts = []string{
	// single raft
	//"192.168.3.129:60062",

	// multiple rafts
	"192.168.3.129:60161",
	"192.168.3.129:60162",
	"192.168.3.129:60163",
}

func GetTestDefaultConnection(hostChange *string) (*sdk.Connection, error) {
	rand.Seed(time.Now().Unix())
	host := hosts[rand.Intn(len(hosts))]

	if hostChange != nil {
		host = *hostChange
	}
	username := "root"
	password := "root"
	crtFile := "./ultipa.crt"
	crtFile = ""
	config := sdk.DefaultConfig{
		//"default", 15, true,
		GraphSetName: "default",
		ResponseWithRequestInfo: true,
		ReadModeNonConsistency: true,
	}
	connect, err := sdk.GetConnection(host, username, password, crtFile, &config)
	if err != nil {
		return nil, err
	}
	return connect, nil
}
func TestLogTitle(str string) {
	log.Println("❗️-------------- TestCase:", str, " --------------")
}
func TestLogSubtitle(str string)  {
	log.Println("⚠️ ******* ", str, " *******")
}
