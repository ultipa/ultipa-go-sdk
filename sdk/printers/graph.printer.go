package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func PrintGraph(graphs []*structs.Graph) {
	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{&simpletable.Cell{Text: "Name"}, &simpletable.Cell{Text: "Description"}, &simpletable.Cell{Text: "Total Node"}, &simpletable.Cell{Text: "Total Edge"}, &simpletable.Cell{Text: "Status"}}
	for _, graph := range graphs {

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			&simpletable.Cell{Text: graph.Name},
			&simpletable.Cell{Text: graph.Description},
			&simpletable.Cell{Text: fmt.Sprint(graph.TotalNodes)},
			&simpletable.Cell{Text: fmt.Sprint(graph.TotalEdges)},
			&simpletable.Cell{Text: graph.Status},
		})

		if len(table.Body.Cells) > 0 {
			table.Println()
		}
	}
}
