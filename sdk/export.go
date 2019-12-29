package sdk

/*
 * Download files from ultipa-server
 */
import (
	"context"
	"fmt"
	"io"

	"github.com/cheggaaa/pb/v3"

	// "io/ioutil"
	"encoding/csv"
	"log"
	"os"

	// "strings"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	// "ultipa-go-sdk/utils"
)

func exportData(client Client, propNames []string, outPath string, _type ultipa.DBType, bar *pb.ProgressBar) {

	var out = ""
	typeName := "NODE"
	headers := []string{"id"}
	if _type == ultipa.DBType_DBNODE {
		out = outPath + "nodes.csv"
	} else if _type == ultipa.DBType_DBEDGE {
		headers = append(headers, "fromID", "toID")
		out = outPath + "edges.csv"
		typeName = "Edge"
	} else {
		log.Println("unknown db type")
		return
	}

	headers = append(headers, propNames...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24*7)
	defer cancel()

	// create folder
	merr := os.MkdirAll(outPath, os.ModePerm)

	if merr != nil {
		fmt.Printf("Create path Failed: %v ", merr)
		return
	}

	// check if file is exist
	info, _ := os.Stat(out)

	if info != nil {
		log.Println("Export file " + out + "nodes.csv or edges.csv is exist, remove it or change the path")
		return
	}

	// open file
	f, e := os.OpenFile(out, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()

	if e != nil {
		fmt.Println(e)
	}

	// create writer
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// start export
	msg, err := client.Export(ctx, &ultipa.ExportRequest{
		DbType:           _type,
		SelectProperties: propNames,
	})

	if err != nil {
		log.Printf("[Error] export error: %v", err)
	}

	log.Print(typeName + " Exporter Started\n")

	//write header
	writer.Write(headers)

	count := 0
	finish := false
	go func() {
		for {

			if finish {
				break
			}

			bar.SetCurrent(int64(count))
			time.Sleep(time.Second)
		}
	}()

	for {
		c, err := msg.Recv()

		if err != nil {
			finish = true
			if err == io.EOF {
				break
			} else {
				log.Println("Failed %v", err)
				break
			}
		}

		if _type == ultipa.DBType_DBEDGE {

			for _, edge := range c.Edges {
				row := []string{edge.Id, edge.FromId, edge.ToId}

				for _, v := range edge.Values {
					row = append(row, v.Value)
				}

				writer.Write(row)
			}
			count += len(c.Edges)
		} else {

			for _, node := range c.Nodes {
				row := []string{node.Id}

				for _, v := range node.Values {
					row = append(row, v.Value)
				}

				writer.Write(row)
			}
			count += len(c.Nodes)
		}

	}

}

func ExportNode(client Client, nodePropNames []string, outPath string) {

	statistic := Statistic(client)

	// create node process

	nodebar := pb.Full.Start(0)
	nodebar.SetTotal(int64(statistic.NodeCount))

	exportData(client, nodePropNames, outPath, ultipa.DBType_DBNODE, nodebar)
	nodebar.SetCurrent(int64(statistic.NodeCount))
	nodebar.Finish()
}

func ExportEdge(client Client, edgePropNames []string, outPath string) {

	statistic := Statistic(client)

	edgebar := pb.Full.Start(0)
	edgebar.SetTotal(int64(statistic.EdgeCount))

	exportData(client, edgePropNames, outPath, ultipa.DBType_DBEDGE, edgebar)
	edgebar.SetCurrent(int64(statistic.EdgeCount))
	edgebar.Finish()
}

func ExportAll(client Client, outPath string) {

	nodePropNames := []string{}

	nodeProperties := GetEdgePropertyInfo(client)
	for _, p := range nodeProperties {
		nodePropNames = append(nodePropNames, p.Name)
	}

	ExportEdge(client, nodePropNames, outPath)

	edgePropNames := []string{}

	edgeProperties := GetEdgePropertyInfo(client)
	for _, p := range edgeProperties {
		edgePropNames = append(edgePropNames, p.Name)
	}

	ExportEdge(client, edgePropNames, outPath)

}
