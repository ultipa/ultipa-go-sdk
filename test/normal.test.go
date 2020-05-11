package main

import "ultipa-go-sdk/sdk"

func GetTestConnect(host string, username string, password string, crt string) (*sdk.Connection, error) {
	connet := sdk.Connection{}
	err := connet.Init(host, username, password, crt)
	if err != nil {
		return nil, err
	}
	return &connet, nil
}
func GetDefaultTestConnect() (*sdk.Connection, error){
	return GetTestConnect("localhost:60061", "root", "root", "./test/ultipa.crt")
}