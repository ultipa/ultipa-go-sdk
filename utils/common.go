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

func BytesToString(bytes []byte) string {
	return string(bytes[:])
}
