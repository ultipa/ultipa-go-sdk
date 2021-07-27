package utils

import (
	"fmt"
	"log"
	"reflect"
)

func Map(slice interface{}, callback func(index int) interface{}) interface{} {
	t := reflect.TypeOf(slice)
	if t.Kind() != reflect.Slice {
		log.Fatalln("utils.Contains -> slice must be a slice")
	}

	s := reflect.ValueOf(slice)
	newSlices := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		newSlices[i] = callback(i)
	}

	return newSlices
}

func IndexOf(slice interface{}, target interface{}) int {
	t := reflect.TypeOf(slice)
	if t.Kind() != reflect.Slice {
		log.Fatalln("utils.Contains -> slice must be a slice")
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if fmt.Sprint(s.Index(i)) == fmt.Sprint(reflect.ValueOf(target)) {
			return i
		}
	}

	return -1
}

func Contains(slice interface{}, target interface{}) bool {

	t := reflect.TypeOf(slice)
	//log.Println(t.Kind())
	if t.Kind() != reflect.Slice {
		log.Fatalln("utils.Contains -> slice must be a slice")
	}

	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		//log.Println(s.Index(i), " || ", target)
		if fmt.Sprint(s.Index(i)) == fmt.Sprint(reflect.ValueOf(target)) {
			return true
		}
	}

	return false
}
