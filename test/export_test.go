package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"testing"
)

func TestExportAsNodesEdges(t *testing.T) {
	//garphName := "Yuri_LegalTech"
	//schemaName := "Judgement"
	//client, _ := GetClient([]string{"192.168.2.142:60062"}, garphName)
	//
	//schema := &structs.Schema{
	//	Name: schemaName,
	//	Properties: []*structs.Property{
	//		{Name: "_id"}, {Name: "_uuid"}, {Name: "link"}, {Name: "type"}, {Name: "decisionDate"}, {Name: "officialDate"}, {Name: "year"}, {Name: "version"}, {Name: "decisionNumber"},
	//		{Name: "ssz"}, {Name: "szDec"}, {Name: "annullamentoConRinvio"}, {Name: "annullamentoSenzaRinvio"}, {Name: "rigettato"}, {Name: "inammissibile"}, {Name: "fullDocument"},
	//	}}

	schemaName := "nodeSchema"
	client, _ := GetClient(hosts, graph)

	schema := &structs.Schema{
		Name: schemaName,
		Properties: []*structs.Property{
			{Name: "_id"}, {Name: "_uuid"}, {Name: "typeInt32"}, {Name: "typeFloat"}, {Name: "typeDouble"}, {Name: "typeInt64"}, {Name: "typeUint32"}, {Name: "typeUint64"}, {Name: "typeDatetime"},
			{Name: "typeString"}, {Name: "typeTimestamp"}, {Name: "typeNotMatch"}, {Name: "typeText"},
		}}
	err := client.ExportAsNodesEdges(schema,
		1000,
		&configuration.RequestConfig{MaxPkgSize: 1024 * 1024 * 100},
		func(nodes []*structs.Node, edges []*structs.Edge) error {
			//printers.PrintNodes(nodes, map[string]*structs.Schema{schemaName: schema})
			t.Log(len(nodes))
			return nil
		})
	if err != nil {
		t.Fatal(err)
	}
}
