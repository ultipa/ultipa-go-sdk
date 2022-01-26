package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/sdk/configuration"
)

func TestLoadConfigYaml(t *testing.T) {
	config, err := configuration.LoadConfigFromYAML("./config.yml")

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("config : %v", config)
}
