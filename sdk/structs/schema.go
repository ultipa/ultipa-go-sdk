package structs

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/utils"
)

type Schema struct {
	Name       string
	Properties []*Property
	Desc       string
	Type       string // node | edge
	DBType     ultipa.DBType
	Total      int
}

func NewSchema(name string) *Schema {
	return &Schema{Name: name, Properties: []*Property{}}
}

func (s *Schema) GetProperty(name string) *Property {
	prop := utils.Find(s.Properties, func(index int) bool { return s.Properties[index].Name == name })

	if prop != nil {
		return prop.(*Property)
	}

	return nil
}

// compare 2 schema is same, or is able to fit schema1 to schema2
func CompareSchemas(schema1 *Schema, schema2 *Schema, fit bool) bool {

	if schema1 == nil {
		return false
	}

	if schema2 == nil {
		return fit
	}

	schema1PropMap := map[string]*Property{}
	schema2PropMap := map[string]*Property{}

	for _, prop := range schema1.Properties {
		schema1PropMap[prop.Name] = prop
	}

	for _, prop := range schema2.Properties {
		schema2PropMap[prop.Name] = prop
	}

	// check one by one
	for name, prop1 := range schema1PropMap {

		prop2 := schema2PropMap[name]

		if fit == true && (prop2 != nil && prop2.Type != prop1.Type) {
			return false
		}

		if fit == false && (prop2 == nil || prop2.Type != prop1.Type) {
			return false
		}
	}

	return true
}
