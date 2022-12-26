package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"ultipa-go-sdk/sdk/structs"
)

func PrintTable(tableData *structs.Table) {
	if tableData == nil {
		fmt.Println("No table data found.")
		return
	}
	table := simpletable.New()

	for _, header := range tableData.GetHeaders() {
		table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: header.Name})
	}

	for _, row := range tableData.GetRows() {
		r := []*simpletable.Cell{}
		for _, field := range *row {
			r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(field)})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Println()
}
