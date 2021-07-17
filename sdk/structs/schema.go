package structs

import (
	"ultipa-go-sdk/sdk/utils"
)

type Schema struct {
	Name string
	Properties []*Property
	Desc string
	Type string // node | edge
	Total int
}

func NewSchema(name string) *Schema {
	return &Schema{ Name: name, Properties: []*Property{} }
}

func (s *Schema) GetProperty(name string)  *Property {
	  prop := utils.Find(s.Properties, func(index int ) bool { return s.Properties[index].Name == name })

	  if prop != nil {
	  	 return prop.(*Property)
	  }

	  return nil
}


