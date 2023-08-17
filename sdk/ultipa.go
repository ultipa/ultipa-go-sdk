//Package sdk provide Ultipa functions to drive ultipa servers
package sdk

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/api"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/connection"
)

var (
	Version = "v4.3.0" //主版本号
)

// New an Ultipa Instance !!!!
func NewUltipa(config *configuration.UltipaConfig) (*api.UltipaAPI, error) {

	config.FillDefault()

	// set connection pool
	pool, err := connection.NewConnectionPool(config)
	if err != nil {
		return nil, err
	}
	// set heartbeat for Connection Pool
	pool.RunHeartBeat()

	if err != nil {
		return nil, err
	}

	return api.NewUltipaAPI(pool), err
}
