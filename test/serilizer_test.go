package test

import (
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/utils"
)

func TestStringAsInterface(t *testing.T) {
	datetime, err := utils.StringAsInterface("1970-01-01", ultipa.PropertyType_DATETIME)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)
}

func TestConvertInterfaceToBytesSafe(t *testing.T) {
	ultipaTime, err := utils.NewTimestampFromString("2010-12-12 00:00:00.000", nil)
	if err != nil {
		t.Fatal(err)
	}
	datetime, err := utils.ConvertInterfaceToBytesSafe(ultipaTime, ultipa.PropertyType_DATETIME)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)

	datetime, err = utils.ConvertInterfaceToBytesSafe("2010-12-12T00:00:00.000Z", ultipa.PropertyType_DATETIME)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)

	ultipaTime, err = utils.NewTimeFromString("1994-12-12 00:00:00")
	if err != nil {
		t.Fatal(err)
	}
	datetime, err = utils.ConvertInterfaceToBytesSafe(ultipaTime, ultipa.PropertyType_DATETIME)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)

	datetime, err = utils.ConvertInterfaceToBytesSafe("1994-12-12 00:00:00", ultipa.PropertyType_DATETIME)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datetime)
}
