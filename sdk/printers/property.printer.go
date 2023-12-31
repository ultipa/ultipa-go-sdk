package printers

import (
	"github.com/alexeyco/simpletable"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"log"
	"strconv"
)

func PrintProperty(properties []*structs.Property) {

	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{&simpletable.Cell{Text: "Name"}, &simpletable.Cell{Text: "Description"}, &simpletable.Cell{Text: "Type"}, &simpletable.Cell{Text: "LTE"}, &simpletable.Cell{Text: "Schema"}}

	for _, prop := range properties {
		propertyTypeStr, err := prop.GetStringType()
		if err != nil {
			log.Panic(err)
		}

		cells := []*simpletable.Cell{
			&simpletable.Cell{Text: prop.Name},
			&simpletable.Cell{Text: prop.Desc},
			&simpletable.Cell{Text: propertyTypeStr},
			&simpletable.Cell{Text: strconv.FormatBool(prop.Lte)},
			&simpletable.Cell{Text: prop.Schema},
		}
		table.Body.Cells = append(table.Body.Cells, cells)
	}
	if len(table.Body.Cells) > 0 {
		table.Println()
	}
}
