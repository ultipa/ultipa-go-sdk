package printers

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
)

func PrintAttr(attr *structs.Attr) {
	if attr == nil {
		fmt.Println("No attr data found.")
		return
	}
	table := simpletable.New()

	table.Header.Cells = []*simpletable.Cell{
		{
			Text: attr.Name,
		},
	}
	stringList := getAttrStr(attr)
	for _, str := range stringList {
		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{
				Text: str,
			},
		})
	}
	table.Println()
}

func getAttrStr(attr *structs.Attr) []string {
	var result []string
	for _, row := range attr.Rows {
		if row == nil {
			result = append(result, "<nil>")
		} else {
			result = append(result, fmt.Sprintf("%v", row))
		}
	}
	return result
}

// getAttrStrWithList @Deprecated
func getAttrStrWithList(attr *structs.Attr) []string {
	switch attr.PropertyType {
	case ultipa.PropertyType_NULL_:
		return []string{}
	case ultipa.PropertyType_LIST:
		var result []string
		for _, row := range attr.Rows {
			result = append(result, getAttrListCellString(row.(*structs.AttrListData)))
		}
		return result
	default:
		var result []string
		for _, row := range attr.Rows {
			result = append(result, fmt.Sprintf("%v", row))
		}
		return result
	}
}

// getAttrStrWithList @Deprecated
func getAttrListCellString(attrListData *structs.AttrListData) string {
	switch attrListData.ResultType {
	case ultipa.ResultType_RESULT_TYPE_NODE:
		return getNodeTableStringWithoutSchema(attrListData.Nodes)
	case ultipa.ResultType_RESULT_TYPE_PATH:
		return getPathTableString(attrListData.Paths)
	case ultipa.ResultType_RESULT_TYPE_EDGE:
		return getEdgeTableStringWithoutSchema(attrListData.Edges)
	case ultipa.ResultType_RESULT_TYPE_ATTR:
		var result []string
		for _, subAttr := range attrListData.Attrs {
			subStrList := getAttrStr(subAttr)
			result = append(result, strings.Join(subStrList, ","))
		}
		return strings.Join(result, ",")
	}
	return ""
}
