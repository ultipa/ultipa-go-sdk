package test

import (
	"log"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestReflectArray(t *testing.T) {

	find := func(list interface{}, it func(index int) bool) interface{} {

		for index := 0; index < reflect.ValueOf(list).Len(); index++ {
			if it(index) {
				return reflect.ValueOf(list).Index(index)
			}
		}

		return nil
	}

	test1 := []string{"a", "b", "c"}

	res1 := find(test1, func(index int) bool {
		return test1[index] == "b"
	})

	log.Println(res1.(string))
}

func TestRegexpMatch(t *testing.T) {
	fnNames := []string{"LTE", "UFE"}
	matcher := regexp.MustCompile(`(\s*|^|\n)(` + strings.Join(fnNames, "|") + `)\(`)

	r := matcher.Match([]byte("LTE().node_property(@User.name)"))
	log.Println(r)
}
