package test

import (
	"testing"
	"time"
	"ultipa-go-sdk/sdk/utils"
)

func TestNewTimeFromString(t *testing.T) {
	str := []string{
		"2022-12-23 19:06:10",
		"69122319:06:10",
		"70122319:06:10",
		"2022/12/23 19:06:10",
		"2022-12-23 19:06:10Z",
		"22-12-23 19:06:10Z",
		"2022-12-23 19:06:10+0800",
		"2022-12-23 19:06:10+0800",
		"2022-12-23T19:06:11+0800",
		"2022122319:06:12+0800",
		"22122319:06:13+0800",
		"02122319:06:14+0800",
		"-1586903608",
		"1999-11-30 00:00:00Z",
	}

	for i, s := range str {
		ultipatime, err := utils.NewDatetimeFromString(s)
		if err != nil {
			t.Fatalf("failed %d:%v", i, err)
		}
		t.Log(ultipatime)
	}
}

func TestTimeToUltipaTime(t *testing.T) {
	ut := utils.TimeToUltipaTime(nil, nil)
	t.Log(ut)
}

func TestTimeZone(t *testing.T) {
	t.Log(time.Now().Zone())
	tt:=time.Unix(0, 0).UTC()
	ut:=utils.TimeToUltipaTime(&tt ,time.UTC)
	t.Log(ut)
	t.Log(ut.Datetime)
}

func TestUint64ToUltipaTime(t *testing.T) {
	ut:=utils.UltipaTime{
		Datetime: 1,
	}
	t.Log(ut)
	gt := ut.Uint64ToTime(uint64(1))
	t.Log(gt)
}