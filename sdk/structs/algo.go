package structs

import (
	"encoding/json"
	"fmt"
)

type AlgoJsonStruct struct {
	Name string
	Version string
	Description string
	Parameters map[string]string
}

type AlgoParam struct {
	Name string
	Desc string
}

type Algo struct {
	Name   string
	Desc string
	Version string
	Params map[string]*AlgoParam
}

func NewAlgo(name string, paramString string) (*Algo, error) {
	algo := &Algo{
		Name:   name,
		Params: map[string]*AlgoParam{},
	}

	algoJsonStruct := AlgoJsonStruct{
		Parameters: map[string]string{},
	}

	err := json.Unmarshal([]byte(paramString), &algoJsonStruct)

	if err != nil {
		return nil, err
	}

	for k,v := range algoJsonStruct.Parameters {
		algo.Params[k] = &AlgoParam{
			Name: k,
			Desc: v,
		}
	}

	algo.Desc = algoJsonStruct.Description
	algo.Version = algoJsonStruct.Version

	return algo, nil
}

func (algo *Algo) ParamsToString() string {
	str := ""
	for k,v := range algo.Params {
		str += fmt.Sprintf("%v : %v\n", k, v.Desc)
	}
	return str
}
