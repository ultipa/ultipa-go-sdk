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
		//fixme: check schema, proeprty by prefix _
		if strings.Contains(res.Name, "nodeSchema") || strings.Contains(res.Name, "edgeSchema"){
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

	default:
		log.Printf("Got UnHandled Alias %v Type %v \n", dataitem.Alias, dataitem.Type)
	}
}
