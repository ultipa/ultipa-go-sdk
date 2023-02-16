package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math"
	"reflect"
	"strconv"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/types"
)

var default_nil_string = string([]byte{0x00})

// Convert Bytes to GoLang Type and return to an interface
func ConvertBytesToInterface(bs []byte, t ultipa.PropertyType, subTypes []ultipa.PropertyType) (interface{}, error) {
	if IsNull(t, bs) {
		return nil, nil
	}
	switch t {
	case ultipa.PropertyType_STRING:
		return AsString(bs), nil
	case ultipa.PropertyType_INT32:
		return AsInt32(bs), nil
	case ultipa.PropertyType_INT64:
		return AsInt64(bs), nil
	case ultipa.PropertyType_UINT32:
		return AsUint32(bs), nil
	case ultipa.PropertyType_UINT64:
		return AsUint64(bs), nil
	case ultipa.PropertyType_FLOAT:
		return AsFloat32(bs), nil
	case ultipa.PropertyType_DOUBLE:
		return AsFloat64(bs), nil
	case ultipa.PropertyType_DATETIME:
		if len(bs) == 0 {
			return NewTime(0), nil
		}
		value := AsUint64(bs)
		return NewTime(value), nil
	case ultipa.PropertyType_TIMESTAMP:
		value := AsUint32(bs)
		if len(bs) == 0 || value == 0 {
			return NewTimeStamp(0), nil
		}
		return NewTimeStamp(int64(value)), nil
	case ultipa.PropertyType_TEXT:
		return AsString(bs), nil
	case ultipa.PropertyType_BLOB:
		return bs, nil
	//case ultipa.PropertyType_POINT:
	//	//TODO
	//case ultipa.PropertyType_DECIMAL:
	//TODO
	case ultipa.PropertyType_LIST:
		return deserializeList(bs, subTypes)
	//	//TODO
	//case ultipa.PropertyType_SET:
	//	//TODO
	//case ultipa.PropertyType_MAP:
	//	//TODO
	//case ultipa.PropertyType_UNSET:
	//	return nil
	default:
		return nil, nil
	}
}

//deserializeList deserialize bs to list
func deserializeList(bs []byte, subTypes []ultipa.PropertyType) (interface{}, error) {
	listData := &ultipa.ListData{}
	if err := proto.Unmarshal(bs, listData); err != nil {
		return nil, err
	}
	var list []interface{}
	if listData.IsNull {
		return list, nil
	}
	if listData.Values != nil {
		for _, value := range listData.Values {
			element, err := ConvertBytesToInterface(value, subTypes[0], nil)
			if err != nil {
				return nil, err
			}
			list = append(list, element)
		}
	}
	return list, nil
}

//ConvertInterfaceToBytesSafe convert value to []byte, if value is nil, will set default value according to PropertyType t
func ConvertInterfaceToBytesSafe(value interface{}, t ultipa.PropertyType, subTypes []ultipa.PropertyType) ([]byte, error) {
	toConvertValue := value
	if toConvertValue == nil {
		switch t {
		case ultipa.PropertyType_SET:
			return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
		case ultipa.PropertyType_MAP:
			return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
		case ultipa.PropertyType_POINT:
			return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
		case ultipa.PropertyType_DECIMAL:
			return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
		default:
			return GetNullBytes(t), nil
		}
	}
	switch t {
	case ultipa.PropertyType_LIST:
		return SerializeListData(value, subTypes)
	case ultipa.PropertyType_POINT:
		return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
	case ultipa.PropertyType_DECIMAL:
		return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
	case ultipa.PropertyType_SET:
		return SerializeSetData(value, subTypes)
	case ultipa.PropertyType_MAP:
		return nil, errors.New(fmt.Sprintf("unsuppoted ultipa.PropertyType [%s]", t))
	case ultipa.PropertyType_DATETIME:
		switch v := value.(type) {
		case UltipaTime:
			return ConvertInterfaceToBytes(v.Datetime)
		case string:
			uTime, err := NewTimeFromString(v)
			if err != nil {
				return nil, err
			}
			return ConvertInterfaceToBytes(uTime.Datetime)
		default:
			return ConvertInterfaceToBytes(value)
		}
	case ultipa.PropertyType_TIMESTAMP:
		switch v := value.(type) {
		case UltipaTime:
			return ConvertInterfaceToBytes(v.GetTimeStamp())
		case string:
			uTime, err := NewTimeFromString(v)
			if err != nil {
				return nil, err
			}
			return ConvertInterfaceToBytes(uTime.GetTimeStamp())
		default:
			return ConvertInterfaceToBytes(value)
		}
	default:
		return ConvertInterfaceToBytes(toConvertValue)
	}
}

func SerializeListData(list interface{}, subTypes []ultipa.PropertyType) ([]byte, error) {
	if subTypes == nil {
		return nil, errors.New("subTypes is nil, unable to serialize list")
	}
	if len(subTypes) == 0 {
		return nil, errors.New("subTypes is not specified, unable to serialize list")
	}
	listData := &ultipa.ListData{}
	if list == nil {
		listData.IsNull = true
		return proto.Marshal(listData)
	}
	vi := reflect.ValueOf(list)
	for index := 0; index < vi.Len(); index++ {
		//TODO if vi.Index(index) is ListValue?
		bs, err := ConvertInterfaceToBytesSafe(vi.Index(index).Interface(), subTypes[0], nil)
		if err != nil {
			return nil, err
		}
		listData.Values = append(listData.Values, bs)
	}
	return proto.Marshal(listData)
}

func SerializeSetData(set interface{}, subTypes []ultipa.PropertyType) ([]byte, error) {
	if subTypes == nil && len(subTypes) == 0 {
		return nil, errors.New("subTypes is nil, unable to serialize SetData")
	}
	if len(subTypes) == 0 {
		return nil, errors.New("subTypes is not specified, unable to serialize SetData")
	}
	if set == nil {
		setData := &ultipa.SetData{}
		setData.IsNull = true
		return proto.Marshal(setData)
	}
	vi := reflect.ValueOf(set)
	setData := &ultipa.SetData{}
	for index := 0; index < vi.Len(); index++ {
		//TODO if vi.Index(index) is ListValue?
		bs, err := ConvertInterfaceToBytesSafe(vi.Index(index).Interface(), subTypes[0], nil)
		if err != nil {
			return nil, err
		}
		setData.Values = append(setData.Values, bs)
	}
	return proto.Marshal(setData)
}

func SerializeMapData(value interface{}, subTypes []ultipa.PropertyType) ([]byte, error) {
	if subTypes == nil && len(subTypes) == 0 {
		return nil, errors.New("subTypes is nil, unable to serialize SetData")
	}
	if len(subTypes) == 0 {
		return nil, errors.New("subTypes is not specified, unable to serialize SetData")
	}
	if value == nil {
		mapData := &ultipa.MapData{}
		mapData.IsNull = true
		return proto.Marshal(mapData)
	}

	switch t := value.(type) {
	case map[interface{}]interface{}:
		toSerializeValue := value.(map[interface{}]interface{})
		var mapValues []*ultipa.MapValue
		for k, v := range toSerializeValue {
			kb, err := ConvertInterfaceToBytesSafe(k, subTypes[0], nil)
			if err != nil {
				return nil, err
			}
			vb, err := ConvertInterfaceToBytesSafe(v, subTypes[0], nil)
			if err != nil {
				return nil, err
			}
			mapValue := &ultipa.MapValue{
				Key:   kb,
				Value: vb,
			}
			mapValues = append(mapValues, mapValue)
		}
		mapData := &ultipa.MapData{Values: mapValues}
		return proto.Marshal(mapData)
	default:
		return nil, errors.New(fmt.Sprintf("value is not a map, but %s, unable to serialize as MapData", t))
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

func GetDefaultNilString(t ultipa.PropertyType) string {

	switch t {
	case ultipa.PropertyType_INT32:
		fallthrough
	case ultipa.PropertyType_INT64:
		fallthrough
	case ultipa.PropertyType_UINT32:
		fallthrough
	case ultipa.PropertyType_UINT64:
		fallthrough
	case ultipa.PropertyType_FLOAT:
		fallthrough
	case ultipa.PropertyType_DOUBLE:
		return "0"
	case ultipa.PropertyType_DATETIME:
		return "1970-01-01"
	case ultipa.PropertyType_TIMESTAMP:
		return "1970-01-01"
	default:
		return ""
	}

}

func GetDefaultNilInterface(t ultipa.PropertyType) interface{} {

	switch t {
	case ultipa.PropertyType_INT32:
		return math.MaxInt32
	case ultipa.PropertyType_INT64:
		return math.MaxInt64
	case ultipa.PropertyType_UINT32:
		return math.MaxUint32
	case ultipa.PropertyType_UINT64:
		return uint64(math.MaxUint64)
	case ultipa.PropertyType_FLOAT:
		return math.MaxFloat32
	case ultipa.PropertyType_DOUBLE:
		return math.MaxFloat64
	case ultipa.PropertyType_DATETIME:
		return uint64(math.MaxUint64)
	case ultipa.PropertyType_TIMESTAMP:
		return math.MaxUint32
	default:
		return default_nil_string
	}
}

func StringAsInterface(str string, t ultipa.PropertyType) (interface{}, error) {

	str = strings.Trim(str, " ")

	if str == "" {
		str = GetDefaultNilString(t)
	}

	switch t {
	case ultipa.PropertyType_INT32:
		v, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(v), err
	case ultipa.PropertyType_INT64:
		return strconv.ParseInt(str, 10, 64)
	case ultipa.PropertyType_UINT32:
		v, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(v), err
	case ultipa.PropertyType_UINT64:
		return strconv.ParseUint(str, 10, 64)
	case ultipa.PropertyType_FLOAT:
		v, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return nil, err
		}
		return float32(v), err
	case ultipa.PropertyType_DOUBLE:
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		return float64(v), err
	case ultipa.PropertyType_DATETIME:
		v, err := NewTimeFromString(str)
		if err != nil {
			return nil, err
		}
		return v.Datetime, err
	case ultipa.PropertyType_TIMESTAMP:
		v, err := NewTimeFromString(str)
		if err != nil {
			return nil, err
		}
		return v.GetTimeStamp(), err
	default:
		return str, nil
	}

	return nil, nil
}

func StringAsUUID(str string) (types.UUID, error) {
	str = strings.Trim(str, " ")
	v, err := strconv.ParseUint(str, 10, 64)
	return v, err
}
