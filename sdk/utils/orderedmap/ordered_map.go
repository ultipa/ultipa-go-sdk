package orderedmap

import (
	"os"
)

type OrderedMap struct {
	Keys []string
	Data map[string]interface{}
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		Data: map[string]interface{}{},
	}
}

func (om *OrderedMap) HasKey(key string) bool {
	return om.Data[key] != nil
}

func (om *OrderedMap) Set(key string, value interface{}) {
	item := om.Get(key)

	if item == nil {
		om.Keys = append(om.Keys, key)
	}

	om.Data[key] = value
}

func (om *OrderedMap) Delete(key string) {
	item := om.Get(key)
	if item == nil {
		return
	}

	index, err := om.GetKeyIndex(key)

	if err == os.ErrNotExist {
		return
	}

	om.Keys = append(om.Keys[:index], om.Keys[index+1:]...)
}

func (om *OrderedMap) GetKeyIndex(key string) (int, error) {
	for index, k := range om.Keys {
		if k == key {
			return index, nil
		}
	}

	return 0, os.ErrNotExist
}

func (om *OrderedMap) Get(key string) interface{} {
	return om.Data[key]
}

func (om *OrderedMap) ForEach(cb func(key string, value interface{}, index int)) {
	for index, key := range om.Keys {
		v := om.Get(key)
		cb(key, v, index)
	}
}
