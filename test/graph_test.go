package test

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils/logger"
	"github.com/ultipa/ultipa-go-sdk/utils"
	"log"
	"testing"
)

func TestShowGraph(t *testing.T) {
	InitCases()
	client, _ := GetClient(hosts, graph)
	res, err := client.ShowGraph(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}

func TestCreateGraph(t *testing.T) {

	client, err := GetClient(hosts, graph)

	if err != nil {
		log.Println(err)
		return
	}

	client.DropGraph(graph, nil)

	client.CreateGraph(&structs.GraphSet{
		Name: graph,
	}, nil)

	client.SetCurrentGraph(graph)

	resp, err := client.Uql("insert().nodes({}).into(@default)", nil)

	if err != nil {
		logger.PrintError(err.Error())
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		logger.PrintError(resp.Status.Message)
	}

}

func TestDropGraph(t *testing.T) {
	client.DropGraph("test_creation", nil)
}

func TestAsGraph(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	resp, _ := client.Uql("show().graph()", nil)
	graphs, err := resp.Alias(http.RESP_GRAPH_KEY).AsGraphInfos()

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintGraphInfo(graphs)
}

func TestCreateGraphIfNotExist(t *testing.T) {

	client, err := GetClient(hosts, graph)

	if err != nil {
		t.Fatalf("failed to connect to server %v", err)
	}

	client.DropGraph(graph, nil)

	_, _, err = client.CreateGraphIfNotExist(&structs.GraphSet{
		Name: graph,
	}, nil)
	if err != nil {
		t.Fatalf("failed to create graph %v", err)
	}
}

func TestUltipaAPI_AlterGraph(t *testing.T) {
	type args struct {
		oldGraphName string
		newGraphName string
		description  string
		config       *configuration.RequestConfig
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Old graph name is empty",
			args: args{
				oldGraphName: "",
				newGraphName: "validNewGraphName",
				description:  "Valid description",
				config:       nil,
			},
			wantErr: true,
		},
		{
			name: "Old graph name contains illegal characters",
			args: args{
				oldGraphName: "1",
				newGraphName: "validNewGraphName",
				description:  "Valid description",
				config:       nil,
			},
			wantErr: true,
		},
		{
			name: "New graph name is empty",
			args: args{
				oldGraphName: "amz",
				newGraphName: "",
				description:  "new description",
				config:       nil,
			},
			wantErr: false,
		},
		{
			name: "Description is empty",
			args: args{
				oldGraphName: "amz",
				newGraphName: "amz1",
				description:  "",
				config:       nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := client.AlterGraph(tt.args.oldGraphName, tt.args.newGraphName, tt.args.description, tt.args.config)
			_ = rsp
			if (tt.wantErr && err == nil) || (!tt.wantErr && err != nil) {
				t.Errorf("AlterGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
