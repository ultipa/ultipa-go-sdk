package printers

import (
	"github.com/alexeyco/simpletable"
	"strconv"
	"ultipa-go-sdk/sdk/structs"
)

func PrintProperty(properties []*structs.Property) {

	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{&simpletable.Cell{Text: "Name"}, &simpletable.Cell{Text: "Description"}, &simpletable.Cell{Text: "Type"}, &simpletable.Cell{Text: "LTE"}, &simpletable.Cell{Text: "Schema"}}

	for _, prop := range properties {
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
}
