package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func FloatStr2IntStr(floatStr string) (string, error) {
	f, e := strconv.ParseFloat(floatStr, 0)
	return fmt.Sprintf("%.0f", f), e
}
func Str2Int(s string) (int64, error) {
	i, e := strconv.ParseInt(s, 10, 0)
	return i, e
}
func Str2Float (floatStr string) (float64, error)  {
	f, e := strconv.ParseFloat(floatStr, 1000000)
	return f, e
}
func BytesToString(bytes []byte) string {
	return string(bytes[:])
}
func StructToJSONBytes(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}
func ToJSONStringPretty(data interface{}) string {
	bs, err := StructToJSONBytes(data)
	if err != nil {
		return "{}"
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, bs, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return "{}"
	}
	return string(prettyJSON.Bytes())
}
func ToJSONString(data interface{})  string{
	bs, err := StructToJSONBytes(data)
	if err != nil {
		return "{}"
	}
	return BytesToString(bs)
}

func InterfacePoint(d interface{}) *interface{} {
	return &d
}