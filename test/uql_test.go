package test

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
)

func TestUQL(t *testing.T) {

	InitCases()

	for _, c := range cases {

		log.Println("Exec : ", c.UQL)

		resp, err := client.UQL(c.UQL, nil)

		if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
			log.Println(resp.Status.Message)
			continue
		}

		if err != nil {
			log.Fatalln("Test UQL Error : ", err)
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
					res, _ := dataitem.AsSchemas()

					for _, schema := range res {
						fmt.Println("Schema Name: ", schema.Name, "(", schema.TotalNodes,"|",schema.TotalEdges ,")")
						fmt.Println("Description: ", schema.Desc)
						table := simpletable.New()
						table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{Text: "Name"}, &simpletable.Cell{Text: "Description"}, &simpletable.Cell{Text: "Type"}, &simpletable.Cell{Text: "lte"})

						for _, prop := range schema.Properties {
							table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
								&simpletable.Cell{Text: prop.Name},
								&simpletable.Cell{Text: prop.Desc},
								&simpletable.Cell{Text: prop.GetStringType()},
							})
						}

						table.Println()
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

			default:
				log.Println("Got UnHandled type", resp.AliasList)
			}
		}

	}
}
