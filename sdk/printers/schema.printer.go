package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"strconv"
	"ultipa-go-sdk/sdk/structs"
)

func PrintSchema(schemas []*structs.Schema) {

	for _, schema := range schemas {
		fmt.Println("Schema Name: ", schema.Name, "(", schema.Total, ")")
		fmt.Println("Description: ", schema.Desc)
		table := simpletable.New()
		table.Header.Cells = []*simpletable.Cell{&simpletable.Cell{Text: "Name"}, &simpletable.Cell{Text: "Description"}, &simpletable.Cell{Text: "Type"}, &simpletable.Cell{Text: "LTE"}, &simpletable.Cell{Text: "Schema"}}
		//table.Footer.Cells = []*simpletable.Cell{&simpletable.Cell{Span: 4, Text: fmt.Sprint("[", schema.Type, "]Schema : "+schema.Name, "(", schema.Total, ")")}}

		for _, prop := range schema.Properties {
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				&simpletable.Cell{Text: prop.Name},
				&simpletable.Cell{Text: prop.Desc},
				&simpletable.Cell{Text: prop.GetStringType()},
				&simpletable.Cell{Text: strconv.FormatBool(prop.Lte)},
				&simpletable.Cell{Text: prop.Schema},
			})
		}

		if len(table.Body.Cells) > 0 {
			table.Println()
		}

		println("-")

	}
}
