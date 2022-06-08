package test

import (
	"log"
	"testing"
	"time"
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

	log.Println(AB&A, AB&B, AB&C, AB&D)
}

func TestTimeZoneTimestamp(t *testing.T) {
	t1 := "2022-05-19T12:00:00Z"
	t2 := "2022-05-19T12:00:00Z"

	tt1, err := time.Parse(time.RFC3339, t1)

	if err != nil {
		log.Fatalln("E1",err)
	}

	tt2, err := time.Parse("2006-01-02T15:04:05Z07:00", t2)

	if err != nil {
		log.Fatalln("E2", err)
	}

	log.Println(tt1.UTC(), tt2.UTC())
}
