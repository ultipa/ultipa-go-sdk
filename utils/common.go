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
func StuctToJSONString(data interface{})  (string, error){
	bs, err := StructToJSONBytes(data)
	if err != nil {
		return "", err
	}
	return BytesToString(bs), nil
}

func BytesToString(bytes []byte) string {
	return string(bytes[:])
}