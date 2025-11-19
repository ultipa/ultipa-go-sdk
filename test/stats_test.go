package test

import (
	"fmt"
	"testing"
)

func TestStats(t *testing.T) {
	stats, err := client.Stats(nil)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(stats)
}
