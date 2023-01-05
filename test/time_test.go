package test

import (
	"testing"
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

	for i, s := range str {
		ultipatime, err := utils.NewTimeFromString(s)
		if err != nil {
			t.Fatalf("failed %d:%v", i, err)
		}
		t.Log(ultipatime)
	}
}

func TestTimeToUltipaTime(t *testing.T) {
	ut := utils.TimeToUltipaTime(nil)
	t.Log(ut)
}
