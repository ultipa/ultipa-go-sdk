package printers

import (
	"log"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/http"
)

func PrintAny(dataitem *http.DataItem) {

	switch dataitem.Type {
	case ultipa.ResultType_RESULT_TYPE_NODE:
		nodes, schemas, _ := dataitem.AsNodes()
		PrintNodes(nodes, schemas)
	case ultipa.ResultType_RESULT_TYPE_TABLE:
		//handle other table
		res, err := dataitem.AsTable()

		// handle schema table
		if strings.Contains(res.Name, "_nodeSchema") || strings.Contains(res.Name, "_edgeSchema"){
			schemas, err := dataitem.AsSchemas()

			if err != nil {
				log.Fatalln(err)
			}

			PrintSchema(schemas)
			return
		}


		if err != nil {
			log.Fatalln(err)
		}

		PrintTable(res)
	case ultipa.ResultType_RESULT_TYPE_PATH:
		paths, err := dataitem.AsPaths()

		if err != nil {
			log.Fatalln(err)
		}

		PrintPaths(paths)
	case ultipa.ResultType_RESULT_TYPE_ATTR:
		attr, err := dataitem.AsAttr()
		if err != nil {
			log.Fatalln(err)
		}

		PrintAttr(attr)

	case ultipa.ResultType_RESULT_TYPE_ARRAY:
		arr, err := dataitem.AsArray()
		if err != nil {
			log.Fatalln(err)
		}

		PrintArray(arr)
	default:
		log.Printf("Printer Got UnHandled Alias %v Type %v \n", dataitem.Alias, dataitem.Type)
	}
}
