package api

import (
	"errors"
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"strings"
)

func (api *UltipaAPI) CreateHDCGraphBySchema(graphName string, nodeSchemas, edgeSchemas []*structs.Schema, update string, hdcName string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	// `hdc.graph.create("social-user-article", {
	//nodes: {User: ["username"], Article: ["title"] },
	//edges: {Post: []},
	//update: "async"
	//}).to("computation-1")`

	nodeSchema := FormatHdcGraphSchemas(nodeSchemas)
	edgeSchema := FormatHdcGraphSchemas(edgeSchemas)
	if nodeSchema == "" || edgeSchema == "" {
		return nil, errors.New("schema name cannot be empth")

	}

	uql := fmt.Sprintf(`hdc.graph.create("%s", {
	nodes: {%s},
	edges: {%s},
	update: "%s",
	query: "query",
    type: "Graph",
    default: true
	}).to("%s")`, graphName, nodeSchema, edgeSchema, update, hdcName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		api.Logger.Log("create hdc graph failed : " + graphName + " " + err.Error())
		return nil, err
	}

	api.Logger.Log("Creating Graph Request OK! - " + graphName)

	return resp, err
}

func FormatHdcGraphSchemas(schemas []*structs.Schema) string {
	if len(schemas) == 0 {
		return `"*": ["*"]`
	}

	var result []string

	for _, schema := range schemas {
		if schema.Name == "" {
			return ""
		}

		if len(schema.Properties) == 0 {
			// 如果 Properties 为空
			result = append(result, fmt.Sprintf(`%s: ["*"]`, schema.Name))
		} else {
			// 拼接每个 Property.Name
			var propertyNames []string
			for _, prop := range schema.Properties {
				propertyNames = append(propertyNames, prop.Name)
			}
			result = append(result, fmt.Sprintf(`%s: ["%s"]`, schema.Name, strings.Join(propertyNames, `", "`)))
		}
	}

	return strings.Join(result, ", ")

}

func (api *UltipaAPI) ShowHDCGraph(requestConfig *configuration.RequestConfig) ([]*structs.Projection, error) {
	//if graphName != "" {
	//	graphName = fmt.Sprintf(`"%s"`, graphName)
	//}
	//
	//uql := fmt.Sprintf(`hdc.graph.show(%s)`, graphName)
	//
	//return api.Uql(uql, requestConfig)
	resp, err := api.Uql("hdc.graph.show()", requestConfig)
	if err != nil {
		return nil, err
	}

	projections, err := resp.Alias(http.RESP_PROJECT_KEY).AsProjections()
	if err != nil {
		return nil, err
	}

	return projections, nil
}

func (api *UltipaAPI) DropHDCGraph(graphName string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`hdc.graph.drop('%s')`, graphName)

	return api.Uql(uql, requestConfig)
}

func (api *UltipaAPI) ShowHDCAlgo(algoName string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	// TODO
	// get algoList from _algoList_from_hdc ?

	uql := fmt.Sprintf(`hdc.server.show('%s')`, algoName)
	return api.Uql(uql, requestConfig)
}
