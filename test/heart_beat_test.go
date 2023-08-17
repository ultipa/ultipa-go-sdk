package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"log"
	"testing"
	"time"
)

func TestHeartBeat(t *testing.T) {

	var err error
	config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts:     hosts,
		Username:  username,
		Password:  password,
		HeartBeat: 1,
		Debug:     true,
	})

	client, err = sdk.NewUltipa(config)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("TestHeartBeat - Sleep")
	time.Sleep(100 * time.Second)

}
