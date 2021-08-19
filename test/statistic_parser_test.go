package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/utils"
)

func TestParseStatistic(t *testing.T) {

	res, err := http.ParseStatistic(&ultipa.Table{
		Headers: []*ultipa.Header{
			{
				PropertyName: "total_time_cost",
				PropertyType: ultipa.PropertyType_STRING,
			},
			{
				PropertyName: "engine_time_cost",
				PropertyType: ultipa.PropertyType_STRING,
			},
		},
		TableRows: []*ultipa.TableRow{
			{
				Values: [][]byte{
					[]byte("10"),
					[]byte("20"),
				},
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	utils.PrintJSON(res)
}
