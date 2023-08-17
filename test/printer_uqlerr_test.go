package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"testing"
)

func TestPrintUQLErr(t *testing.T) {
	c := "[3-5]aaa find\nExported"
	printers.PrintUqlErr(c)
}
