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
	pool , err := connection.NewConnectionPool(config)

	// set heartbeat for Connection Pool
	pool.RunHeartBeat()

	if  err != nil {
		return nil ,err
	}

	return api.NewUltipaAPI(pool), err
}