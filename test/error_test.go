package test

import (
	"reflect"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

func TestErrorType(t *testing.T) {
	err := utils.NewLeaderNotYetElectedError("")
	if reflect.TypeOf(err).Elem().String() != "utils.LeaderNotYetElectedError" {
		t.Error("not instance of utils.LeaderNotYetElectedError")
	} else {
		t.Log("ok")
	}
}
