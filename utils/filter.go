package utils

import (
	// "google.golang.org/grpc"
	"encoding/json"
	"fmt"
	"reflect"
	// "strings"
	// "ultipa-go-sdk/pkg"
	"github.com/robertkrimen/otto"

	ultipa "ultipa-go-sdk/rpc"

	"strconv"
)

// message Filter {
//   enum OPTION{
//     FILTER_OPT_AND = 0;
//     FILTER_OPT_OR = 1;
//   }
//   OPTION filter_option = 1;
//   repeated FilterValue values = 2;
//  }
//  message FilterValue {
//   string column_name =  1;
//   string left_value  =  2;
//   string right_value =  3;
// }

type FilterCondition struct {
	Key   string
	Left  string
	Right string
}

type Filter = ultipa.Filter

// type Filter struct {
// 	FilterOption string
// 	Conditions   []FilterCondition
// }

func NewFilterCondition(key string, operator string, value []string) []FilterCondition {
	var conditions []FilterCondition
	var condition FilterCondition
	condition.Key = key
	const MaxInt = int(^uint32(0) >> 1)

	// fmt.Println(MaxInt)
	switch operator {
	case "=":
		condition.Left = value[0]
		condition.Right = value[0]
		conditions = append(conditions, condition)
		break
	case ">":
		condition.Left = value[0]
		condition.Right = strconv.Itoa(MaxInt)
		conditions = append(conditions, condition)
		break
	case "<":
		condition.Left = strconv.Itoa(-MaxInt - 1)
		condition.Right = value[0]
		conditions = append(conditions, condition)
		break
	case "<>":
		condition.Left = value[0]
		condition.Right = value[1]
		conditions = append(conditions, condition)
		break
	case "in":
		for _, v := range value {
			cond := NewFilterCondition(key, "=", []string{v})
			conditions = append(conditions, cond[0])
		}
		break
	}
	return conditions
}

// NewFilter return Filter for filter things, AND | OR,
func NewFilter(option string, conditions []FilterCondition) ultipa.Filter {

	var uOption ultipa.Filter_OPTION
	var uConditions []*ultipa.FilterValue

	switch option {
	case "AND":
		uOption = ultipa.Filter_FILTER_OPT_AND
		break
	case "OR":
		uOption = ultipa.Filter_FILTER_OPT_OR
		break
	}

	for _, c := range conditions {
		uConditions = append(uConditions, &ultipa.FilterValue{PropertyName: c.Key, LeftValue: c.Left, RightValue: c.Right})
	}

	return ultipa.Filter{
		FilterOption: uOption,
		Values:       uConditions,
	}
}

// StringToFilters return filters from a string like {"name":"ultipa"}
func StringToFilters(str string) []FilterCondition {

	var filters []FilterCondition

	if len(str) < 1 {
		return filters
	}

	vm := otto.New()

	value, _ := vm.Run("JSON.stringify(" + str + ")")
	{
		// value is an int64 with a value of 16
		value, _ := value.ToString()
		var dat map[string]interface{}
		j := []byte(value)

		json.Unmarshal(j, &dat)

		// fmt.Printf("%#v \n", dat)
		for k, v := range dat {

			arr := reflect.ValueOf(v)
			values := []string{}
			var filter []FilterCondition

			// fmt.Println(arr.Kind())
			if arr.Kind() == reflect.Slice {
				vArr := v.([]interface{})

				for _, v2 := range vArr {
					values = append(values, fmt.Sprint(v2))
					// fmt.Println("values", values)
				}

				filter = NewFilterCondition(k, "<>", values)
			} else {

				values = append(values, fmt.Sprint(v))
				filter = NewFilterCondition(k, "=", values)

			}

			filters = append(filters, filter...)
		}

	}

	// fmt.Printf(" %v \n", filters)

	return filters
}

func allTOString(item interface{}) {

}
