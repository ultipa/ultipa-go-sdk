package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func PrintNodes(nodes []*structs.Node, schemas map[string]*structs.Schema) {
	if len(nodes) == 0 {
		fmt.Println("No node data found.")
		return
	}
	fmt.Println(getNodeTableString(nodes, schemas))
}

func PrintNodesWithoutSchema(nodes []*structs.Node) {
	if len(nodes) == 0 {
		fmt.Println("No node data found.")
		return
	}
	fmt.Println(getNodeTableStringWithoutSchema(nodes))
}

func getNodeTableStringWithoutSchema(nodes []*structs.Node) string {
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
	var schemaMap = make(map[string]*structs.Schema)
	for schemaName, propertyList := range schemaPropertiesMap {
		schema := structs.NewSchema(schemaName)
		schema.DBType = ultipa.DBType_DBNODE
		for _, propertyName := range propertyList {
			schema.Properties = append(schema.Properties, &structs.Property{
				Name:   propertyName,
				Schema: schemaName,
			})
		}
		schemaMap[schemaName] = schema
	}
	return getNodeTableString(nodes, schemaMap)
}

func getNodeTableString(nodes []*structs.Node, schemas map[string]*structs.Schema) string {
	var lastSchema string
	var table *simpletable.Table
	switchSchema := false
	for _, node := range nodes {
		schema := schemas[node.Schema]
		if node.Schema != lastSchema {
			switchSchema = true
			lastSchema = node.Schema
		} else {
			switchSchema = false
		}

		if table != nil && switchSchema {
			fmt.Println(table.String())
			table = nil
			switchSchema = false
		}
		if table == nil {
			table = simpletable.New()
			table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: "ID"}, &simpletable.Cell{Align: simpletable.AlignCenter, Text: "UUID"}, &simpletable.Cell{Align: simpletable.AlignCenter, Text: "Schema"})
			for _, prop := range schema.Properties {
				table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: prop.Name})
			}
		}
		r := []*simpletable.Cell{
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: node.GetID()},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(node.GetUUID())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(node.GetSchema())},
		}

		for i := 3; i < len(table.Header.Cells); i++ {

			headerKey := table.Header.Cells[i].Text
			vv := node.Values.Get(headerKey)
			r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(vv)})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	if table != nil {
		return table.String()
	}
	return ""
}
