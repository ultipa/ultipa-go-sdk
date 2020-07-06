package test

import (
	"fmt"
	"testing"
)

type Abc struct {
	Status string
}
type E struct {
	*Abc
	Message string
}

func TestStruct(t *testing.T) {

	var a = E{
		Message: "123",
		Abc: &Abc{Status: "b123"},
	}

	fmt.Println(a)
	fmt.Println(a.Status,a.Message)
}

