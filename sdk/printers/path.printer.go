package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"ultipa-go-sdk/sdk/structs"
)

func PrintPaths(paths []*structs.Path) {
	table := simpletable.New()

	table.Header.Cells = []*simpletable.Cell{
		{Text: "#"},
		{Text: "Path"},
	}

	for num, path := range paths {
		row := []*simpletable.Cell{{Text: fmt.Sprint(num)}}
		pathString := ""
		for index, edge := range path.GetEdges() {
			node := path.GetNodes()[index]
			d1 := "-"
			d2 := "-"

			if edge.GetFrom() == node.GetID() {
				d2 = "->"
			} else {
				d1 = "<-"
			}

			pathString = fmt.Sprintf("(%v) %v [%v] %v ", node.ID, d1, edge.UUID, d2)
		}
		pathString += fmt.Sprintf("(%v)", path.GetLastNode().GetID())
		row = append(row, &simpletable.Cell{Text: pathString})
		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.Println()
}