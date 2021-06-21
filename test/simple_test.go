package test

import (
	"log"
	"reflect"
	"testing"
)

func TestReflectArray(t *testing.T) {

	find := func (list interface{}, it func(index int) bool) interface {} {

		for index := 0; index < reflect.ValueOf(list).Len(); index++ {
			if it(index) {
				return reflect.ValueOf(list).Index(index)
			}
		}

		return nil
	}

	test1 := []string{"a","b","c"}

	res1 := find(test1, func(index int) bool {
		return test1[index] == "b"
	})

	log.Println(res1.(string))
}
