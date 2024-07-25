package printers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func PrintSchema(schemas []*structs.Schema) {

	for _, schema := range schemas {
		fmt.Println("Schema Name: ", schema.Name, "(", schema.Total, ")")
		fmt.Println("Description: ", schema.Desc)
		table := simpletable.New()
		table.Header.Cells = []*simpletable.Cell{{Text: "Name"}, {Text: "Description"}, {Text: "Type"}, {Text: "LTE"}, {Text: "Schema"}}
		//table.Footer.Cells = []*simpletable.Cell{&simpletable.Cell{Span: 4, Text: fmt.Sprint("[", schema.Type, "]Schema : "+schema.Name, "(", schema.Total, ")")}}

		for _, prop := range schema.Properties {

			propertyTypeStr, err := prop.GetStringType()
			if err != nil {
				log.Panic(err)
			}

			rowCells := []*simpletable.Cell{
				{Text: prop.Name},
				{Text: prop.Desc},
				{Text: propertyTypeStr},
				{Text: strconv.FormatBool(prop.Lte)},
				{Text: prop.Schema},
			}

			table.Body.Cells = append(table.Body.Cells, rowCells)
		}

		if len(table.Body.Cells) > 0 {
			table.Println()
		}

		println("-")

	}
}
