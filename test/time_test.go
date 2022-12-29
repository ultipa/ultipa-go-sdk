package test

import (
	"testing"
	"time"
	"ultipa-go-sdk/sdk/utils"
)

func TestNewTimeFromString(t *testing.T) {
	str := []string{
		"2022-12-23 19:06:10",
		"2022-12-23T19:06:11+0800",
		"2022122319:06:12+0800",
		"22122319:06:13+0800",
		"2122319:06:14+0800",
	}
	//utils.SetTimestampPrintingLocation("Europe/Paris")
	for i, s := range str {
		ultipatime, err := utils.NewTimeFromString(s)
		if err != nil {
			t.Fatalf("failed %d:%v", i, err)
		}
		t.Log(ultipatime)
	}
}

func TestLoadLocation(t *testing.T) {
	_, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Fatal(err)
	}
}
