package test

import (
	"log"
	"testing"
)

func TestShowTask(t *testing.T) {
	tasks, err := client.ShowTask("", 0, nil)
	if err != nil {
		t.Fatalf("show task error %v", err)
	}

	for _, task := range tasks {
		log.Println(task)
	}
}
