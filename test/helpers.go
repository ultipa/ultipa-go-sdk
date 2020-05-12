package test

import (
	"log"
	"ultipa-go-sdk/sdk"
)

func GetTestDefaultConnection(hostChange *string) (*sdk.Connection, error) {
	host := "192.168.3.129:60162"
	if hostChange != nil {
		host = *hostChange
	}
	username := "root"
	password := "root"
	//crtFile := "./test/ultipa.crt" // debug
	crtFile := "./ultipa.crt"
	connect, err := sdk.GetConnection(host, username, password, crtFile)
	if err != nil {
		return  nil, err
	}
	return  connect, nil
}
func Debug(format string, args ...interface{})  {
	log.Printf(format, args...)
}