package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"reflect"
	"testing"
)

func TestErrorType(t *testing.T) {
	err := utils.NewLeaderNotYetElectedError("")
	if reflect.TypeOf(err).Elem().String() != "utils.LeaderNotYetElectedError" {
		t.Error("not instance of utils.LeaderNotYetElectedError")
	} else {
		t.Log("ok")
	}
}
