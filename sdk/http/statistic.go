package http

import (
	"reflect"
	"strconv"
	ultipa "ultipa-go-sdk/rpc"
)

// 存储返回的统计信息
type Statistic struct {
	RowAffected int `key:"rows_affected" type:"int"`
	TotalCost   int `key:"total_time_cost" type:"int"`
	EngineCost  int `key:"engine_time_cost" type:"int"`
}

func ParseStatistic(table *ultipa.Table) (*Statistic, error) {

	stat := Statistic{}

	if table == nil { return &stat , nil }

	kv := map[string]string{}

	for index, header := range table.Headers {

		key := header.PropertyName
		value := table.TableRows[0].Values[index]
		kv[key] = string(value)
	}

	statReflect := reflect.TypeOf(stat)
	statReflectValues := reflect.ValueOf(&stat)

	for i := 0; i < statReflect.NumField(); i++ {
		field := statReflect.Field(i)
		fieldValue := statReflectValues.Elem().Field(i)

		key := field.Tag.Get("key")
		t := field.Tag.Get("type")

		v := kv[key]

		if v == "" {
			continue
		}

		switch t {
		case "int":
			vv, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return &stat, err
			}
			fieldValue.SetInt(vv)
		}
	}

	return &stat, nil
}
