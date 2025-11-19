package test

import (
	"fmt"
	"testing"
)

func TestShowTask(t *testing.T) {
	tasks, err := client.ShowTask("", 0, nil)
	if err != nil {
		t.Errorf("show task error %v", err)
	}
	fmt.Printf("%v\n", tasks[0])

}
