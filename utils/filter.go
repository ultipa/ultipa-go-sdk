package utils

import (
	// "google.golang.org/grpc"
	// "fmt"
	// "ultipa-go-sdk/pkg"
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

type filterCondition struct {
	Key   string
	Left  string
	Right string
}

type filter struct {
	FilterOption string
	Conditions   []filterCondition
}

func NewFilterCondition(key string, operator string, value []string) []filterCondition {
	var conditions []filterCondition
	var condition filterCondition
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
func NewFilter(option string, conditions []filterCondition) ultipa.Filter {

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
