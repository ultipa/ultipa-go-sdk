package test

import (
	"github.com/joho/godotenv"
	"github.com/ultipa/ultipa-go-sdk/sdk"
	"github.com/ultipa/ultipa-go-sdk/sdk/api"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"log"
	"strings"
	"testing"
)

var env map[string]string
var client *api.UltipaAPI
var hosts []string
var username string
var password string
var graph string

func TestMain(m *testing.M) {
	var err error
	env, err = godotenv.Read(".env")

	if err != nil {
		log.Fatalln(err)
	}

	hosts = strings.Split(env["hosts"], ",")
	username, password, graph = env["username"], env["password"], env["graph"]

	client, err = GetClient(hosts, graph)

	if err != nil {
		log.Fatalln(err)
	}

	m.Run()
}

func TestPing(t *testing.T) {
	client, _ = GetClient(hosts, graph)
	client.Test()
}

func GetClient(hosts []string, graphName string) (*api.UltipaAPI, error) {
	var err error
	config, err := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts:        hosts,
		Username:     username,
		Password:     password,
		DefaultGraph: graphName,
		Debug:        true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	client, err = sdk.NewUltipa(config)

	if err != nil {
		log.Fatalln(err)
	}

	return client, err
}
