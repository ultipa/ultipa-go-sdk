package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

type DataItem struct {
	Alias string
	Type  ultipa.ResultType
	Data  interface{}
}

func NewDataItem() *DataItem {
	return &DataItem{}
}

func NodeTableToNodes(nt *ultipa.NodeTable, alias string) ([]*structs.Node, map[string]*structs.Schema) {

	schemas := map[string]*structs.Schema{}
	nodes := []*structs.Node{}

	for _, oSchema := range nt.Schemas {
		schema := structs.NewSchema(oSchema.SchemaName)
		schemas[schema.Name] = schema
		for _, header := range oSchema.Properties {
			schema.Properties = append(schema.Properties, &structs.Property{Name: header.PropertyName, Type: header.PropertyType})
		}
	}

	for _, oNode := range nt.NodeRows {
		node := &structs.Node{
			Name:   alias,
			ID:     oNode.Id,
			UUID:   oNode.Uuid,
			Schema: oNode.SchemaName,
		}

		// set values
		node.Values = structs.NewValues()
		schema := schemas[oNode.SchemaName]
		for index, v := range oNode.Values {
			prop := schema.Properties[index]
			node.Values.Set(prop.Name, utils.ConvertBytesToInterface(v, prop.Type))
		}

		nodes = append(nodes, node)
	}

	return nodes, schemas
}

func EdgeTableToEdges(et *ultipa.EdgeTable, alias string) ([]*structs.Edge, map[string]*structs.Schema) {

	schemas := map[string]*structs.Schema{}
	edges := []*structs.Edge{}

	for _, oSchema := range et.Schemas {
		schema := structs.NewSchema(oSchema.SchemaName)
		schemas[schema.Name] = schema
		for _, header := range oSchema.Properties {
			schema.Properties = append(schema.Properties, &structs.Property{Name: header.PropertyName, Type: header.PropertyType})
		}
	}

	for _, oEdge := range et.EdgeRows {
		edge := &structs.Edge{
			Name:     alias,
			UUID:     oEdge.Uuid,
			From:     oEdge.FromId,
			FromUUID: oEdge.Uuid,
			To:       oEdge.ToId,
			ToUUID:   oEdge.ToUuid,
			Schema:   oEdge.SchemaName,
		}

		// set values
		edge.Values = structs.NewValues()
		schema := schemas[oEdge.SchemaName]
		for index, v := range oEdge.Values {
			prop := schema.Properties[index]
			edge.Values.Set(prop.Name, utils.ConvertBytesToInterface(v, prop.Type))
		}

		edges = append(edges, edge)

	}
	return edges, schemas
}

func (di *DataItem) AsNodes() (nodes []*structs.Node, schemas map[string]*structs.Schema, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return nodes, schemas, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_NODE {
		return nil, schemas, errors.New("DataItem " + di.Alias + " is not Type Node")
	}

	oNodes := di.Data.(*ultipa.NodeAlias)

	nodes, schemas = NodeTableToNodes(oNodes.NodeTable, oNodes.Alias)

	return nodes, schemas, nil
}

func (di *DataItem) AsEdges() (edges []*structs.Edge, schemas map[string]*structs.Schema, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return edges, schemas, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_EDGE {
		return nil, schemas, errors.New("DataItem " + di.Alias + " is not Type Edge")
	}

	oEdges := di.Data.(*ultipa.EdgeAlias)

	edges, schemas = EdgeTableToEdges(oEdges.EdgeTable, oEdges.Alias)

	return edges, schemas, nil
}

func (di *DataItem) AsPaths() (paths []*structs.Path, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return paths, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_PATH {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Paths")
	}

	oPaths := di.Data.(*ultipa.PathAlias)

	for _, oPath := range oPaths.Paths {

		path := structs.NewPath()
		path.Name = oPaths.Alias
		path.Nodes, path.NodeSchemas = NodeTableToNodes(oPath.NodeTable, path.Name)
		path.Edges, path.EdgeSchemas = EdgeTableToEdges(oPath.EdgeTable, path.Name)

		paths = append(paths, path)

	}

	return paths, err
}

func (di *DataItem) AsTable() (table *structs.Table, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return table, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Table")
	}

	table = structs.NewTable()

	oTable := di.Data.(*ultipa.Table)

	table.Name = oTable.TableName

	for _, header := range oTable.Headers {
		h := &structs.Property{
			Name: header.PropertyName,
			Type: header.PropertyType,
		}
		table.Headers = append(table.Headers, h)
	}

	for _, row := range oTable.TableRows {

		r := structs.Row{}

		for index, field := range row.Values {
			r = append(r, utils.ConvertBytesToInterface(field, table.Headers[index].Type))
		}

		table.Rows = append(table.Rows, &r)
	}

	return table, err
}

//find().nodes() as nodes group by nodes.year as y return y,collect(nodes._id)
func (di *DataItem) AsArray() (arr *structs.Array, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return arr, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ARRAY {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Array")
	}

	arr = structs.NewArray()

	oArray := di.Data.(*ultipa.ArrayAlias)

	arr.Name = oArray.Alias

	for _, oRow := range oArray.Elements {
		r := structs.Row{}

		for _, field := range oRow.Values {
			r = append(r, utils.ConvertBytesToInterface(field, oArray.PropertyType))
		}

		arr.Rows = append(arr.Rows, &r)
	}

	return arr, err
}

func (di *DataItem) AsAttr() (attr *structs.Attr, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return attr, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Attribute list")
	}

	attr = structs.NewAttr()

	oAttr := di.Data.(*ultipa.AttrAlias)

	attr.Name = oAttr.Alias

	for _, v := range oAttr.Values {
		attr.Rows = append(attr.Rows, utils.ConvertBytesToInterface(v, oAttr.PropertyType))
	}

	return attr, err
}

// the types will be tables and alias is nodeSchema and edgeSchema
func (di *DataItem) AsGraphs() (graphs []*structs.Graph, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return graphs, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " should be a table as pre-condition")
	}

	table := di.Data.(*ultipa.Table)

	if table.TableName != RESP_GRAPH_KEY {
		return nil, errors.New("DataItem " + di.Alias + " is not a Graph list")
	}

	for _, row := range table.TableRows {
		//0:id, 1: name, 2: totalNodes ,3:totalEdges
		values := row.GetValues()
		graph := structs.Graph{}
		graph.ID = string(values[0])
		graph.Name = string(values[1])
		graph.TotalNodes, _ = utils.Str2Uint64(utils.AsString(values[2]))
		graph.TotalEdges, _ = utils.Str2Uint64(utils.AsString(values[3]))

		graphs = append(graphs, &graph)
	}

	return graphs, err
}

// the types will be tables and alias is nodeSchema and edgeSchema
func (di *DataItem) AsSchemas() (schemas []*structs.Schema, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return schemas, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " should be a table as pre-condition")
	}

	table := di.Data.(*ultipa.Table)

	if table.TableName != RESP_NODE_SCHEMA_KEY && table.TableName != RESP_EDGE_SCHEMA_KEY {
		return nil, errors.New("DataItem " + di.Alias + " is not a Schema list")
	}

	// node | edge
	Type := ""
	// store index to get total number
	TotalIndex := 0
	switch table.TableName {
	case RESP_NODE_SCHEMA_KEY:
		Type = "node"
		TotalIndex = 3
	case RESP_EDGE_SCHEMA_KEY:
		Type = "edge"
		TotalIndex = 4
	}

	for _, row := range table.TableRows {
		//0:name, 1: description, 2: json(properties),3:totalNodes, 4:totalEdges
		values := row.GetValues()
		schema := structs.NewSchema(string(values[0]))
		schema.Desc = string(values[1])
		schema.Type = Type
		schema.Total, _ = strconv.Atoi(utils.AsString(values[TotalIndex]))
		propertyJson := values[2]

		schema.DBType, err = structs.GetDBTypeByString(schema.Type)

		if err != nil {
			return nil, err
		}

		var props []*struct {
			Name        string
			Type        string
			description string
			lte         bool
		}

		err = json.Unmarshal(propertyJson, &props)

		if err != nil {
			log.Fatalln(err)
		}

		for _, prop := range props {
			p := structs.Property{
				Name: prop.Name,
				Desc: prop.description,
			}
			p.SetTypeByString(prop.Type)
			schema.Properties = append(schema.Properties, &p)
		}

		schemas = append(schemas, schema)
	}

	return schemas, err
}

func (di *DataItem) AsAlgos() ([]*structs.Algo, error) {

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " should be a table(algo) as pre-condition")
	}

	table, err := di.AsTable()

	if err != nil {
		return nil, err
	}

	var algos []*structs.Algo

	algoDatas := table.ToKV()

	for _, algoData := range algoDatas {

		algo, err := structs.NewAlgo(algoData.Data["name"].(string), algoData.Data["param"].(string))


		if err != nil {
			return nil, errors.New(fmt.Sprint(err.Error(), algoData))
		}

		algos = append(algos, algo)
	}

	return algos, nil
}
