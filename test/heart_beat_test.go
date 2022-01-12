package test

import (
	"log"
	"testing"
	"time"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/configuration"
)

func TestHeartBeat(t *testing.T) {

	var err error
	config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts:    []string{"192.168.1.86:63101"},
		Username: "root",
		Password: "root",
		HeartBeat: 1,
	})

	client, err = sdk.NewUltipa(config)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("TestHeartBeat - Sleep")
	time.Sleep(100 * time.Second)

}
