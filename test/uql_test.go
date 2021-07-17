package test

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

func TestUQL(t *testing.T) {

	InitCases()

	for _, c := range cases {

		log.Println("Exec : ", c.UQL)

		//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
		resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})

		if err != nil {
			log.Fatalln(err)
		}

		if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
			log.Println(resp.Status.Message)
			continue
		}

		for _, a := range c.Alias {

			dataitem := resp.Alias(a)
			//t := c.Type[index]

			switch dataitem.Type {
			case ultipa.ResultType_RESULT_TYPE_NODE:

				nodes, schemas, _ := dataitem.AsNodes()

				var lastSchema string
				var table *simpletable.Table
				for _, node := range nodes {
					schema := schemas[node.Schema]
					if node.Schema != lastSchema {

						if table != nil {
							fmt.Println(table.String())
						}
						table = simpletable.New()
						lastSchema = node.Schema
						table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: "ID"}, &simpletable.Cell{Align: simpletable.AlignCenter, Text: "UUID"}, &simpletable.Cell{Align: simpletable.AlignCenter, Text: "Schema"})
						for _, prop := range schema.Properties {
							table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: prop.Name})
						}
					}

					r := []*simpletable.Cell{
						&simpletable.Cell{Align: simpletable.AlignCenter, Text: node.GetID()},
						&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(node.GetUUID())},
						&simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(node.GetSchema())},
					}

					node.Values.ForEach(func(v interface{}, key string) error {
						r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(v)})
						return nil
					}, nil)

					table.Body.Cells = append(table.Body.Cells, r)
				}

				if table != nil {
					table.Println()
				}
			case ultipa.ResultType_RESULT_TYPE_TABLE:

				// handle schema table
				if c.Type == "schema" {
					res, err := dataitem.AsSchemas()

					if err != nil {
						log.Fatalln(err)
					}

					for _, schema := range res {
						fmt.Println("Schema Name: ", schema.Name, "(", schema.Total, ")")
						fmt.Println("Description: ", schema.Desc)
						table := simpletable.New()
						table.Header.Cells = []*simpletable.Cell{&simpletable.Cell{Text: "Name"}, &simpletable.Cell{Text: "Description"}, &simpletable.Cell{Text: "Type"}, &simpletable.Cell{Text: "LTE"}}
						//table.Footer.Cells = []*simpletable.Cell{&simpletable.Cell{Span: 4, Text: fmt.Sprint("[", schema.Type, "]Schema : "+schema.Name, "(", schema.Total, ")")}}

						for _, prop := range schema.Properties {
							table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
								&simpletable.Cell{Text: prop.Name},
								&simpletable.Cell{Text: prop.Desc},
								&simpletable.Cell{Text: prop.GetStringType()},
								&simpletable.Cell{},
							})
						}

						if len(table.Body.Cells) > 0 {
							table.Println()
						}

						println("-")

					}

					continue
				}

				//handle other table
				res, err := dataitem.AsTable()
				if err != nil {
					log.Fatalln(err)
				}

				table := simpletable.New()

				for _, header := range res.GetHeaders() {
					table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Align: simpletable.AlignCenter, Text: header.Name})
				}

				for _, row := range res.GetRows() {
					r := []*simpletable.Cell{}
					for _, field := range *row {
						r = append(r, &simpletable.Cell{Align: simpletable.AlignCenter, Text: fmt.Sprint(field)})
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				table.Println()
			case ultipa.ResultType_RESULT_TYPE_PATH:
				paths, err := dataitem.AsPaths()

				if err != nil {
					log.Fatalln(err)
				}

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

			default:
				log.Println("Got UnHandled type", resp.AliasList)
			}
		}

	}
}
