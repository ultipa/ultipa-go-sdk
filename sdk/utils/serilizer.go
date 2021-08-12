package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
	ultipa "ultipa-go-sdk/rpc"
)

// Convert Bytes to GoLang Type and return to an interface
func ConvertBytesToInterface(bs []byte, t ultipa.UltipaPropertyType) interface{} {
	switch t {
	case ultipa.UltipaPropertyType_STRING:
		return AsString(bs)
	case ultipa.UltipaPropertyType_INT32:
		return AsInt32(bs)
	case ultipa.UltipaPropertyType_INT64:
		return AsInt64(bs)
	case ultipa.UltipaPropertyType_UINT32:
		return AsUint32(bs)
	case ultipa.UltipaPropertyType_UINT64:
		return AsUint64(bs)
	case ultipa.UltipaPropertyType_FLOAT:
		return AsFloat32(bs)
	case ultipa.UltipaPropertyType_DOUBLE:
		return AsFloat64(bs)
	case ultipa.UltipaPropertyType_DATETIME:
		return NewTime(AsUint64(bs))
	case ultipa.UltipaPropertyType_UNSET:
		return nil
	default:
		return nil

	}
}

func ConvertInterfaceToBytes(value interface{}) ([]byte, error) {
	v := []byte{}

	switch t := value.(type) {
	case int32:
		v = make([]byte, 4)
		binary.BigEndian.PutUint32(v, uint32(value.(int32)))
	case int:
		v = make([]byte, 4)
		binary.BigEndian.PutUint32(v, uint32(value.(int32)))
	case string:
		v = []byte(value.(string))
	case int64:
		v = make([]byte, 8)
		binary.BigEndian.PutUint64(v, uint64(value.(int64)))
	case uint32:
		v = make([]byte, 4)
		binary.BigEndian.PutUint32(v, uint32(value.(uint32)))
	case uint64:
		v = make([]byte, 8)
		binary.BigEndian.PutUint64(v, uint64(value.(uint64)))
	case float32:
		v = make([]byte, 4)
		binary.BigEndian.PutUint32(v, math.Float32bits(value.(float32)))
	case float64:
		v = make([]byte, 8)
		binary.BigEndian.PutUint64(v, math.Float64bits(value.(float64)))
	default:
		return nil, errors.New(fmt.Sprint("not supported ultipa type : ", t))
	}

	return v, nil
}

func readBytes(value []byte, out interface{}) {
	bs := make([]byte, len(value))
	copy(bs, value)
	buff := bytes.NewBuffer(bs)
	binary.Read(buff, binary.BigEndian, out)
}

func AsInt32(value []byte) int32 {
	var out int32
	readBytes(value, &out)
	return out
}

func AsInt64(value []byte) int64 {
	var out int64
	readBytes(value, &out)
	return out
}

func AsUint32(value []byte) uint32 {
	var out uint32
	readBytes(value, &out)
	return out
}

func AsUint64(value []byte) uint64 {
	var out uint64
	readBytes(value, &out)
	return out
}

func AsFloat32(value []byte) float32 {
	var out float32
	readBytes(value, &out)
	return out
}

func AsFloat64(value []byte) float64 {
	var out float64
	readBytes(value, &out)
	return out
}

func AsString(value []byte) string {
	return string(value)
}

func AsBool(value []byte) bool {
	var out uint16
	readBytes(value, &out)
	if out == 1 {
		return true
	} else {
		return false
	}
}

func StringAsInterface(str string, t ultipa.UltipaPropertyType) (interface{}, error) {

	switch t {
	case ultipa.UltipaPropertyType_INT32:
		v, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(v), err
	case ultipa.UltipaPropertyType_INT64:
		return strconv.ParseInt(str, 10, 32)
	default:
		return str, nil
	}

	return nil, nil
}
