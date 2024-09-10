package test

import (
	"log"
	"testing"
)

func TestTop(t *testing.T) {
	client.SetCurrentGraph("go_sdk_test")
	tops, err := client.Top(nil)
	if err != nil {
		t.Fatalf("exec top error, %v", err)
	}

	for _, top := range tops {
		log.Println(top)
	}
}
