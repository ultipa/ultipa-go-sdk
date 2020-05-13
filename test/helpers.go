package test

import (
	"log"
	"ultipa-go-sdk/sdk"
)

func GetTestDefaultConnection(hostChange *string) (*sdk.Connection, error) {
	//host := "192.168.3.129:60062"
	host := "192.168.3.171:60061"
	//host = "192.168.3.185:60061" // listUser has data
	if hostChange != nil {
		host = *hostChange
	}
	username := "root"
	password := "root"
	crtFile := "./ultipa.crt"
	connect, err := sdk.GetConnection(host, username, password, crtFile)
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
