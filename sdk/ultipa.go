package sdk

import (
	"ultipa-go-sdk/sdk/api"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/connection"
)

// New an Ultipa Instance !!!!
func NewUltipa(config *configuration.UltipaConfig) (*api.UltipaAPI, error){

	config.FillDefault()

	// set connection pool
	pool := connection.NewConnectionPool(config)

	return api.NewUltipaAPI(pool), nil
}