package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
)

func PrintAttr(attr *structs.Attr) {
	if attr == nil {
		fmt.Println("No attr data found.")
		return
	}
	table := attrToTable(attr)

	table.Println()
}

func attrToTable(attr *structs.Attr) *simpletable.Table {
	table := simpletable.New()

	table.Header.Cells = []*simpletable.Cell{
		{
			Text: attr.Name,
		},
	}

	switch attr.PropertyType {
	case ultipa.PropertyType_SET:
		fallthrough
	case ultipa.PropertyType_LIST:
		//TODO
		//listDataRows := attr.Rows
		//for _, row := range listDataRows {
		//	listDataRow := row.(*structs.AttrListData)
		//
		//}
	case ultipa.PropertyType_MAP:
		mapDataRows := attr.Rows
		for _, row := range mapDataRows {
			mapDataRow := row.(*structs.AttrMapData)
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{
					Text: fmt.Sprintf("%v:%v", mapDataRow.Key.Rows, mapDataRow.Value.Rows),
				},
			})
		}
	default:
		for _, row := range attr.Rows {

			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{
					Text: fmt.Sprint(row),
				},
			})
		}
	}
	return table
}
