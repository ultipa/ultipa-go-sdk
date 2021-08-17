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

func TestBitIsInclude(t *testing.T) {

	const (
		U uint32 = 0
		A uint32 = 1
		B uint32 = 2
		C uint32 = 4
		D uint32 = 8
	)

	AB := A | B

	log.Println(AB & A, AB & B, AB & C, AB & D)
}