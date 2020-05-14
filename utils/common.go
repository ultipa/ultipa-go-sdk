package utils

import (
	"encoding/json"
	"fmt"
)

//StructToJSONBytes 把数据转成json的bytes
func StructToJSONBytes(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}
func StructToJSONString(data interface{})  (string, error){
	bs, err := StructToJSONBytes(data)
	if err != nil {
		return "", err
	}
	return BytesToString(bs), nil
}
func StructToPrettyJSONString(data interface{})  (string, error)  {
	b, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return BytesToString(b), nil
}

func BytesToString(bytes []byte) string {
	return string(bytes[:])
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}