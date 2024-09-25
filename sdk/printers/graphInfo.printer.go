package printers

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func PrintGraphSet(graphs []*structs.GraphSet) {
	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{
		{Text: "ID"},
		{Text: "Name"},
		{Text: "Status"},
		{Text: "Description"},
		{Text: "Total Node"},
		{Text: "Total Edge"},
		{Text: "Shards"},
		{Text: "Slot Num"},
		{Text: "Replica Num"},
	}

	for _, graph := range graphs {
		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Text: fmt.Sprint(graph.ID)}, // 添加 ID
			{Text: graph.Name},
			{Text: graph.Status},
			{Text: graph.Description},
			{Text: fmt.Sprint(graph.TotalNodes)},
			{Text: fmt.Sprint(graph.TotalEdges)},
			{Text: graph.Shards},     // 添加 Shards
			{Text: graph.SlotNum},    // 添加 SlotNum
			{Text: graph.ReplicaNum}, // 添加 ReplicaNum
		})
	}

	if len(table.Body.Cells) > 0 {
		table.Println()
	}
}
