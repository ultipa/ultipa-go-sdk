package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"ultipa-go-sdk/sdk/structs"
)

func PrintEdges(edges []*structs.Edge, schemas map[string]*structs.Schema) {
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
		}
		if table == nil {
			table = simpletable.New()
			table.Header.Cells = append(table.Header.Cells,
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "UUID"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "FROM_UUID"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "TO_UUID"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "SCHEMA"})
			for _, prop := range schema.Properties {
				table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: prop.Name})
			}
		}

		r := []*simpletable.Cell{
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetUUID())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetFrom())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetTo())},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(edge.GetSchema())},
		}

		for i := 4; i < len(table.Header.Cells); i++ {

			headerKey := table.Header.Cells[i].Text
			vv := edge.Values.Get(headerKey)
			r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(vv)})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	if table != nil {
		table.Println()
	}
}
