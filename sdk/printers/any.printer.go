package printers

import (
	"fmt"
	"log"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/http"
	//"ultipa-go-sdk/sdk/http"
)

func PrintAny(dataitem *http.DataItem) {
	if dataitem == nil {
		fmt.Println("No dataItem found.")
		return
	}
	switch dataitem.Type {
	case ultipa.ResultType_RESULT_TYPE_NODE:
		nodes, schemas, _ := dataitem.AsNodes()
		PrintNodes(nodes, schemas)
	case ultipa.ResultType_RESULT_TYPE_EDGE:
		edges, schemas, _ := dataitem.AsEdges()
		PrintEdges(edges, schemas)
	case ultipa.ResultType_RESULT_TYPE_TABLE:
		//handle other table
		res, err := dataitem.AsTable()

		// handle schema table
		if strings.Contains(res.Name, http.RESP_NODE_SCHEMA_KEY) || strings.Contains(res.Name, http.RESP_EDGE_SCHEMA_KEY) {
			schemas, err := dataitem.AsSchemas()

			if err != nil {
				log.Fatalln(err)
			}

			PrintSchema(schemas)
			return
		}

		// handle algo table
		if strings.Contains(res.Name, http.RESP_ALGOS_KEY) {
			algos, err := dataitem.AsAlgos()

			if err != nil {
				log.Fatalln(err)
			}

			PrintAlgoList(algos)
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

	//case ultipa.ResultType_RESULT_TYPE_ARRAY:
	//	arr, err := dataitem.AsArray()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	PrintArray(arr)
	default:
		log.Printf("Printer Got UnHandled Alias %v Type %v \n", dataitem.Alias, dataitem.Type)
	}
}
