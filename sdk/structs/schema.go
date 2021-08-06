package structs

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/utils"
)

type Schema struct {
	Name string
	Properties []*Property
	Desc string
	Type string // node | edge
	DBType ultipa.DBType
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


// todo:
// compare 2 schema is same, or is able to fit a to b
func CompareSchemas(schema1 *Schema, schema2 *Schema, fit bool) bool {

	return true
}