package test

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	TestLogTitle("Connect test")
	connet, err := GetTestDefaultConnection(nil)
	if err != nil {
		t.Error(err)
	}
	result, err1 := connet.TestConnect(nil)
	if err1 != nil {
		t.Error(err1)
	}
	if result == false {
		t.Error(result)
	}
	fmt.Printf("test connect -> %v\n", result)
}
