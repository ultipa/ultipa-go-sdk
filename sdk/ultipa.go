//Package sdk provide Ultipa functions to drive ultipa servers
package sdk

import (
	"ultipa-go-sdk/sdk/api"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/connection"
)

var (
	Version  = "v4.2.2" //主版本号
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