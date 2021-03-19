package test

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"regexp"
	"testing"
	"time"
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
		Abc:     &Abc{Status: "b123"},
	}

	fmt.Println(a)
	fmt.Println(a.Status, a.Message)
}

func TestBigEndianFloat(t *testing.T) {

	var op uint32
	var f float32
	v := make([]byte, 4)
	v3 := make([]byte, 4)

	f = 20.5
	op = 1 << 31

	a := op ^ math.Float32bits(f)

	binary.BigEndian.PutUint32(v, math.Float32bits(f))
	binary.BigEndian.PutUint32(v3, a)

	log.Printf("%08b\n", v)
	log.Printf("%08b\n", v3)

}

func TestStringToFilter(t *testing.T) {
	// step one replace all the strings as __STR_<INDEX>
	str := `@amz_node.name == "zh\"ang"`
	matcher := regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)
	index := 0
	str = matcher.ReplaceAllStringFunc(str, func(s string) string {
		log.Println(s)
		index++
		return fmt.Sprint("__STR_", index)
	})

	log.Println(str)
}

func TestNumberCompare(t *testing.T) {
	a := 20.5
	b := 20

	log.Println(int(a), a > float64(b))
}

func TestChannel(t *testing.T) {
	stop := make(chan bool)

	go func(chan bool) {
		stop <- false
		time.Sleep(time.Second)
		stop <- false
		time.Sleep(time.Second)
		stop <- false
		time.Sleep(time.Second)
		stop <- true
	}(stop)

	for {
		var a bool
		select {
		case a = <-stop:
			if a == false {
				log.Println("not yet")
			}
			if a {
				return
			}
		}
	}
}

func TestReferenceOfSlice(t *testing.T) {

	f := func(s *[]string) {
		*s = append(*s, "a")
	}

	ss := []string{}
	f(&ss)

	log.Println(ss)
}

func TestMultiChannel(t *testing.T) {
	var i chan int

	go func(c chan int) {
		for {
			select {
			case ii := <-c:
				log.Println("routine 1 : ", ii)
			}
		}
	}(i)

	go func(c chan int) {
		for {
			select {
			case ii := <-c:
				log.Println("routine 2 : ", ii)
			}
		}
	}(i)

	for {
		time.Sleep(time.Second)
		i <- 1
	}
}
