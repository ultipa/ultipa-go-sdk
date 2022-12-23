package utils

import "reflect"

func Find(list interface{}, it func(index int) bool) interface{} {
	vi := reflect.ValueOf(list)
	if vi.IsNil() {
		return nil
	}
	for index := 0; index < vi.Len(); index++ {
		if it(index) {
			return reflect.ValueOf(list).Index(index).Interface()
		}
	}

	return nil
}
