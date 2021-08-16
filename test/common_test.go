package test

import (
	"log"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{2}

	deleteIndex := 0
	a = append(a[:deleteIndex], a[deleteIndex+1:]...)

	log.Fatalln(a)
}


