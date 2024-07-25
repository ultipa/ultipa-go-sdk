package printers

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func PrintGraphSet(graphs []*structs.GraphSet) {
	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{{Text: "Name"}, {Text: "Description"}, {Text: "Total Node"}, {Text: "Total Edge"}, {Text: "Status"}}
	for _, graph := range graphs {

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Text: graph.Name},
			{Text: graph.Description},
			{Text: fmt.Sprint(graph.TotalNodes)},
			{Text: fmt.Sprint(graph.TotalEdges)},
			{Text: graph.Status},
		})

		if len(table.Body.Cells) > 0 {
			table.Println()
		}
	}
}
