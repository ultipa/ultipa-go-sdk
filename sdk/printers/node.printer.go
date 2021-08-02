package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"ultipa-go-sdk/sdk/structs"
)

func PrintNodes(nodes []*structs.Node, schemas map[string]*structs.Schema) {
	var lastSchema string
	var table *simpletable.Table
	for _, node := range nodes {
		schema := schemas[node.Schema]
		if node.Schema != lastSchema {

			if table != nil {
				fmt.Println(table.String())
			}
			table = simpletable.New()
			lastSchema = node.Schema
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

		node.Values.ForEach(func(v interface{}, key string) error {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(v)})
			return nil
		}, nil)

		table.Body.Cells = append(table.Body.Cells, r)
	}

	if table != nil {
		table.Println()
	}
}
