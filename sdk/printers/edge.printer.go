package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func PrintEdges(edges []*structs.Edge, schemas map[string]*structs.Schema) {
	if len(edges) == 0 {
		fmt.Println("No edge data found.")
		return
	}
	fmt.Println(getEdgeTableString(edges, schemas))
}

func PrintEdgesWithoutSchema(edges []*structs.Edge) {
	if len(edges) == 0 {
		fmt.Println("No edge data found.")
		return
	}
	fmt.Println(getEdgeTableStringWithoutSchema(edges))
}

func getEdgeTableStringWithoutSchema(edges []*structs.Edge) string {
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
	var schemaMap = make(map[string]*structs.Schema)
	for schemaName, propertyList := range schemaPropertiesMap {
		schema := structs.NewSchema(schemaName)
		schema.DBType = ultipa.DBType_DBEDGE
		for _, propertyName := range propertyList {
			schema.Properties = append(schema.Properties, &structs.Property{
				Name:   propertyName,
				Schema: schemaName,
			})
		}
		schemaMap[schemaName] = schema
	}
	return getEdgeTableString(edges, schemaMap)
}

func getEdgeTableString(edges []*structs.Edge, schemas map[string]*structs.Schema) string {
	var lastSchema string
	var table *simpletable.Table
	switchSchema := false
	for _, edge := range edges {
		schema := schemas[edge.Schema]
		if edge.Schema != lastSchema {
			switchSchema = true
			lastSchema = edge.Schema
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
			table.Header.Cells = append(table.Header.Cells,
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "UUID"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "FROM_UUID"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "FROM"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "TO_UUID"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "TO"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "SCHEMA"})
			for _, prop := range schema.Properties {
				table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: prop.Name})
			}
		}

		r := []*simpletable.Cell{
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetUUID())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.FromUUID)},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetFrom())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.ToUUID)},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetTo())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetSchema())},
		}

		for i := 6; i < len(table.Header.Cells); i++ {

			headerKey := table.Header.Cells[i].Text
			vv := edge.Values.Get(headerKey)
			r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(vv)})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	if table != nil {
		return table.String()
	}
	return ""
}
