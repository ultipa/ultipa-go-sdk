package test

import (
	"log"
	"testing"
	"time"

	"github.com/ultipa/ultipa-go-sdk/sdk"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
)

func TestHeartBeat(t *testing.T) {

	var err error
	config, err := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts:     hosts,
		Username:  username,
		Password:  password,
		HeartBeat: 1,
		//Debug:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	client, err = sdk.NewUltipa(config)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("TestHeartBeat - Sleep")
	time.Sleep(10 * time.Second)

}
