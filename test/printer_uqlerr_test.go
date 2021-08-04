package test

import (
	"testing"
	"ultipa-go-sdk/sdk/printers"
)

func TestPrintUQLErr(t *testing.T) {
	c := "[3-5]aaa find\nExported"
	printers.PrintUqlErr(c)
}
