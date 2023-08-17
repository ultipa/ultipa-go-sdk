package structs

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/types"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

type Edge struct {
	Name     string
	From     types.ID
	To       types.ID
	FromUUID types.UUID
	ToUUID   types.UUID
	UUID     types.UUID
	Schema   string
	Values   *Values
}

func NewEdge() *Edge {
	return &Edge{
		Values: NewValues(),
	}
}

func NewEdgeFromMetaData(md *MetaData) *Edge {
	return &Edge{
		From:   md.From,
		To:     md.To,
		UUID:   md.UUID,
		Schema: md.Schema,
		Values: md.Values,
	}
}

func NewEdgeFromEdgeRow(schema *Schema, edgeRow *ultipa.EntityRow) (*Edge, error) {
	newEdge := NewEdge()

	newEdge.UUID = edgeRow.Uuid
	newEdge.From = edgeRow.FromId
	newEdge.To = edgeRow.ToId
	newEdge.FromUUID = edgeRow.FromUuid
	newEdge.ToUUID = edgeRow.ToUuid
	newEdge.Name = edgeRow.SchemaName

	for index, v := range edgeRow.GetValues() {
		prop := schema.Properties[index]
		value, err := utils.ConvertBytesToInterface(v, prop.Type, prop.SubTypes)
		if err != nil {
			return nil, err
		}
		newEdge.Values.Set(prop.Name, value)
	}

	return newEdge, nil
}

func (edge *Edge) GetUUID() types.UUID {
	return edge.UUID
}

func (edge *Edge) GetFrom() types.ID {
	return edge.From
}

func (edge *Edge) GetTo() types.ID {
	return edge.To
}

func (edge *Edge) GetSchema() string {
	return edge.Schema
}

func (edge *Edge) GetValues() *Values {
	return edge.Values
}

// get a value by key
func (edge *Edge) Get(key string) interface{} {
	return edge.Values.Get(key)
}

// get a value by key
func (edge *Edge) GetBytes(key string) ([]byte, error) {
	v := edge.Values.Get(key)
	return utils.ConvertInterfaceToBytes(v)
}

// GetBytesSafe get []byte value by key, if value is nil, then return default value of PropertyType t
func (edge *Edge) GetBytesSafe(key string, t ultipa.PropertyType, subTypes []ultipa.PropertyType, req *configuration.RequestConfig) ([]byte, error) {
	v := edge.Values.Get(key)
	return utils.ConvertInterfaceToBytesSafe(v, t, subTypes, req)
}

// set a value by key
func (edge *Edge) Set(key string, value interface{}) error {

	edge.Values.Set(key, value)
	return nil
}

func ConvertStringEdges(schema *Schema, edges []*Edge, req *configuration.RequestConfig) {
	// Obtain the configured time zone information
	// timezoneOffset > timeZone
	location := utils.GetLocationFromConfig(req)

	// For by Schema, not nodes value
	for _, edge := range edges {
		for _, prop := range schema.Properties {
			stri := edge.Values.Get(prop.Name)

			str := ""
			if stri == nil {
				str = utils.GetDefaultNilString(prop.Type)
			} else {
				str = stri.(string)
			}

			v, err := utils.StringAsInterface(str, prop.Type, location)

			if err != nil {
				continue
			}
			edge.Values.Set(prop.Name, v)
		}
	}
}

func GetSchemasOfEdgeList(edges []*Edge) map[string]*Schema {
	var schemaPropertiesMap = make(map[string][]string)
	for _, edge := range edges {
		propertyList, ok := schemaPropertiesMap[edge.Schema]
		if !ok {
			schemaPropertiesMap[edge.Schema] = []string{}
		}
		for property, _ := range edge.Values.Data {
			if !utils.Contains(propertyList, property) {
				propertyList = append(propertyList, property)
				schemaPropertiesMap[edge.Schema] = propertyList
			}
		}
	}
	var schemaMap = make(map[string]*Schema)
	for schemaName, propertyList := range schemaPropertiesMap {
		schema := NewSchema(schemaName)
		schema.DBType = ultipa.DBType_DBEDGE
		for _, propertyName := range propertyList {
			schema.Properties = append(schema.Properties, &Property{
				Name:   propertyName,
				Schema: schemaName,
			})
		}
		schemaMap[schemaName] = schema
	}
	return schemaMap
}
