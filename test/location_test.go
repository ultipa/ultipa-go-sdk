package test

import (
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/utils"
)

func TestLocation(t *testing.T) {
	config1 := &configuration.RequestConfig{
		TimezoneOffset: 3600,
	}
	location1 := utils.GetLocationFromConfig(config1)
	timestamp1, _ := utils.StringAsInterface("1970-01-01 00:00:00", ultipa.PropertyType_TIMESTAMP, location1)
	t.Log(timestamp1)

	config2 := &configuration.RequestConfig{
		Timezone: "+0300",
	}
	location2 := utils.GetLocationFromConfig(config2)
	timestamp2, _ := utils.StringAsInterface("1970-01-01 00:00:00", ultipa.PropertyType_TIMESTAMP, location2)
	t.Log(timestamp2)

	config3 := &configuration.RequestConfig{
		Timezone: "Europe/Paris",
	}
	location3 := utils.GetLocationFromConfig(config3)
	timestamp3, _ := utils.StringAsInterface("2038-01-19 03:14:07", ultipa.PropertyType_TIMESTAMP, location3)
	t.Log(timestamp3)

	config4 := &configuration.RequestConfig{
		Timezone: "Asia/Shanghai",
	}
	location4 := utils.GetLocationFromConfig(config4)
	timestamp4, _ := utils.StringAsInterface("2038-01-19 03:14:07", ultipa.PropertyType_TIMESTAMP, location4)
	t.Log(timestamp4)
}
