package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"ultipa-go-sdk/sdk/structs"
)

func PrintAttr(attr *structs.Attr) {
	table := simpletable.New()

	table.Header.Cells = []*simpletable.Cell{
		{
			Text: attr.Name,
		},
	}

	for _, row := range attr.Rows {

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{
				Text: fmt.Sprint(row),
			},
		})
	}

	table.Println()
}