package test

import (
	"fmt"
	"testing"
)

func TestTop(t *testing.T) {
	tops, err := client.Top(nil)
	if err != nil {
		t.Error(err)
	}

	for _, top := range tops {
		fmt.Println(top)
	}
}
