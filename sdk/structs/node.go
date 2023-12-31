package structs

import (
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/types"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

type Node struct {
	Name   string
	ID     types.ID
	UUID   types.UUID
	Schema string
	Values *Values
}

func NewNode() *Node {
	return &Node{Values: NewValues()}
}

func (node *Node) GetID() types.ID {
	return node.ID
}

func (node *Node) GetUUID() types.UUID {
	return node.UUID
}

func (node *Node) GetSchema() string {
	return node.Schema
}

func (node *Node) GetValues() *Values {
	return node.Values
}

// get a value by key
func (node *Node) Get(key string) interface{} {
	return node.Values.Get(key)
}

// get a value by key
func (node *Node) GetBytes(key string) ([]byte, error) {
	v := node.Values.Get(key)
	return utils.ConvertInterfaceToBytes(v)
}

// GetBytesSafe get []byte value by key, if value is nil, then return default value of PropertyType t
func (node *Node) GetBytesSafe(key string, t ultipa.PropertyType, subTypes []ultipa.PropertyType, req *configuration.RequestConfig) ([]byte, error) {
	v := node.Values.Get(key)
	return utils.ConvertInterfaceToBytesSafe(v, t, subTypes, req)
}

// set a value by key
func (node *Node) Set(key string, value interface{}) error {

	//todo: check value type
	node.Values.Set(key, value)
	return nil
}

func (node *Node) UpdateByValueID() {
	id := node.Get("_id")
	//uuid := node.Get("_uuid")
	if id != nil {
		node.ID = id.(string)
	}
}

func NewNodeFromNodeRow(schema *Schema, nodeRow *ultipa.EntityRow) (*Node, error) {
	newNode := NewNode()

	newNode.ID = nodeRow.Id
	newNode.UUID = nodeRow.Uuid
	newNode.Name = nodeRow.SchemaName
	newNode.Schema = schema.Name
	for index, v := range nodeRow.GetValues() {
		prop := schema.Properties[index]
		value, err := utils.ConvertBytesToInterface(v, prop.Type, prop.SubTypes)
		if err != nil {
			return nil, err
		}
		newNode.Values.Set(prop.Name, value)
	}
	return newNode, nil
}

func ConvertStringNodes(schema *Schema, nodes []*Node, req *configuration.RequestConfig) {
	// Obtain the configured time zone information
	// timezoneOffset > timeZone
	location := utils.GetLocationFromConfig(req)

	// For by Schema, not nodes value
	for _, node := range nodes {
		for _, prop := range schema.Properties {
			stri := node.Values.Get(prop.Name)

			str := ""
			if stri == nil {
				str = utils.GetDefaultNilString(prop.Type)
			} else {
				if strtmp, ok := stri.(string); ok {
					str = strtmp
				} else {
					str = fmt.Sprint(str)
				}
			}

			v, err := utils.StringAsInterface(str, prop.Type, location)

			if err != nil {
				continue
			}
			node.Values.Set(prop.Name, v)
		}
	}
}

func GetSchemasOfNodeList(nodes []*Node) map[string]*Schema {
	var schemaPropertiesMap = make(map[string][]string)
	for _, node := range nodes {
		propertyList, ok := schemaPropertiesMap[node.Schema]
		if !ok {
			schemaPropertiesMap[node.Schema] = []string{}
		}
		for property, _ := range node.Values.Data {
			if !utils.Contains(propertyList, property) {
				propertyList = append(propertyList, property)
				schemaPropertiesMap[node.Schema] = propertyList
			}
		}
	}
	var schemaMap = make(map[string]*Schema)
	for schemaName, propertyList := range schemaPropertiesMap {
		schema := NewSchema(schemaName)
		schema.DBType = ultipa.DBType_DBNODE
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
