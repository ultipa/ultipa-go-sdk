package printers

import (
	"github.com/alexeyco/simpletable"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func PrintAlgoList(algos []*structs.Algo) {
	table := simpletable.New()

	table.Header.Cells = []*simpletable.Cell{
		{
			Text: "Algo Name",
		},
		{
			Text: "Version",
		},
		{
			Text: "Description",
		},
		{
			Text: "Parameters",
		},
	}

	for _, algo := range algos {

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{
				Text: algo.Name,
			},
			{
				Text: algo.Desc,
			},
			{
				Text: algo.Version,
			},
			{
				Text: algo.ParamsToString(),
			},
		})
	}

	table.Println()
}
