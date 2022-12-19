package test

import (
	"testing"
	ultipa "ultipa-go-sdk/rpc"
)

func TestAuthenticate(t *testing.T) {
	connection, _ := GetClient([]string{"192.168.1.60:60061"}, "default")

	resp, err := connection.Authenticate(ultipa.AuthenticateType_PERMISSION_TYPE_UQL, "show().graph()", nil)

	if err != nil {
		t.Fatal(err)
	}
	if ultipa.ErrorCode_SUCCESS != resp.Status.GetErrorCode() {
		t.Error("authentication is failed")
	}
}
