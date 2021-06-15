package utils

import "reflect"

func Find(list interface{}, it func(index int) bool) interface{} {

	for index := 0; index < reflect.ValueOf(list).Len(); index++ {
		if it(index) {
			 reflect.ValueOf(list).Index(index)
		}
	}

	return nil
}
