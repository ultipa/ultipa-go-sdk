package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
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

func NodeTableToNodes(nt *ultipa.EntityTable, alias string) ([]*structs.Node, map[string]*structs.Schema, error) {

	schemas := map[string]*structs.Schema{}
	nodes := []*structs.Node{}

	for _, oSchema := range nt.Schemas {
		schema := structs.NewSchema(oSchema.SchemaName)
		schemas[schema.Name] = schema
		for _, header := range oSchema.Properties {
			schema.Properties = append(schema.Properties, &structs.Property{Name: header.PropertyName, Type: header.PropertyType, SubTypes: header.SubTypes})
		}
	}

	for _, oNode := range nt.EntityRows {
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
			value, err := utils.ConvertBytesToInterface(v, prop.Type, prop.SubTypes)
			if err != nil {
				return nil, nil, err
			}
			node.Values.Set(prop.Name, value)
		}

		nodes = append(nodes, node)
	}

	return nodes, schemas, nil
}

func EdgeTableToEdges(et *ultipa.EntityTable, alias string) ([]*structs.Edge, map[string]*structs.Schema, error) {

	schemas := map[string]*structs.Schema{}
	edges := []*structs.Edge{}

	for _, oSchema := range et.Schemas {
		schema := structs.NewSchema(oSchema.SchemaName)
		schemas[schema.Name] = schema
		for _, header := range oSchema.Properties {
			schema.Properties = append(schema.Properties, &structs.Property{Name: header.PropertyName, Type: header.PropertyType})
		}
	}

	for _, oEdge := range et.EntityRows {
		edge := &structs.Edge{
			Name:     alias,
			UUID:     oEdge.Uuid,
			From:     oEdge.FromId,
			FromUUID: oEdge.FromUuid,
			To:       oEdge.ToId,
			ToUUID:   oEdge.ToUuid,
			Schema:   oEdge.SchemaName,
		}

		// set values
		edge.Values = structs.NewValues()
		schema := schemas[oEdge.SchemaName]
		for index, v := range oEdge.Values {
			prop := schema.Properties[index]
			value, err := utils.ConvertBytesToInterface(v, prop.Type, prop.SubTypes)
			if err != nil {
				return nil, nil, err
			}
			edge.Values.Set(prop.Name, value)
		}

		edges = append(edges, edge)

	}
	return edges, schemas, nil
}

func (di *DataItem) AsNodes() (nodes []*structs.Node, schemas map[string]*structs.Schema, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return nodes, schemas, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_NODE {
		return nil, schemas, errors.New("DataItem " + di.Alias + " is not Type Node")
	}
	if di.Data == nil {
		return nil, nil, nil
	}
	oNodes := di.Data.(*ultipa.NodeAlias)

	return NodeTableToNodes(oNodes.NodeTable, oNodes.Alias)
}

func (di *DataItem) AsEdges() (edges []*structs.Edge, schemas map[string]*structs.Schema, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return edges, schemas, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_EDGE {
		return nil, schemas, errors.New("DataItem " + di.Alias + " is not Type Edge")
	}

	if di.Data == nil {
		return nil, nil, nil
	}

	oEdges := di.Data.(*ultipa.EdgeAlias)

	return EdgeTableToEdges(oEdges.EdgeTable, oEdges.Alias)
}

func (di *DataItem) AsPaths() (paths []*structs.Path, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return paths, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_PATH {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Path")
	}

	if di.Data == nil {
		return nil, nil
	}

	pathAlias := di.Data.(*ultipa.PathAlias)

	return parsePaths(pathAlias.Paths, pathAlias.Alias)
}

func parsePaths(oPaths []*ultipa.Path, name string) (paths []*structs.Path, err error) {
	for _, oPath := range oPaths {
		path := structs.NewPath()
		path.Name = name
		path.Nodes, path.NodeSchemas, err = NodeTableToNodes(oPath.NodeTable, path.Name)
		if err != nil {
			return nil, err
		}
		path.Edges, path.EdgeSchemas, err = EdgeTableToEdges(oPath.EdgeTable, path.Name)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func (di *DataItem) AsTable() (table *structs.Table, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return table, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Table")
	}

	oTable := di.Data.(*ultipa.Table)

	table = structs.NewTable()
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
			value, err := utils.ConvertBytesToInterface(field, table.Headers[index].Type, table.Headers[index].SubTypes)
			if err != nil {
				return nil, err
			}
			r = append(r, value)
		}

		table.Rows = append(table.Rows, &r)
	}

	return table, err
}

//AsArray find().nodes() as nodes group by nodes.year as y return y,collect(nodes._id)
//func (di *DataItem) AsArray() (arr *structs.Array, err error) {
//
//	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
//		return arr, nil
//	}
//
//	if di.Type != ultipa.ResultType_RESULT_TYPE_ARRAY {
//		return nil, errors.New("DataItem " + di.Alias + " is not Type Array")
//	}
//
//	arr = structs.NewArray()
//
//	oArray := di.Data.(*ultipa.ArrayAlias)
//
//	arr.Name = oArray.Alias
//
//	for _, oRow := range oArray.Elements {
//		r := structs.Row{}
//
//		for _, field := range oRow.Values {
//			//TODO, check has subTypes or not?
//			value, err := utils.ConvertBytesToInterface(field, oArray.PropertyType, nil)
//			if err != nil {
//				return nil, err
//			}
//			r = append(r, value)
//		}
//
//		arr.Rows = append(arr.Rows, &r)
//	}
//
//	return arr, err
//}

func (di *DataItem) AsAttr() (attr *structs.Attr, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return attr, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Attribute list")
	}

	attrAlias := di.Data.(*ultipa.AttrAlias)
	oAttr := attrAlias.Attr

	return parseAttr(oAttr, attrAlias.Alias)
}

func parseAttr(oAttr *ultipa.Attr, name string) (*structs.Attr, error) {
	attr := structs.NewAttr()
	attr.Name = name

	if oAttr != nil {
		attr.PropertyType = oAttr.PropertyType
		switch oAttr.PropertyType {
		case ultipa.PropertyType_SET:
			fallthrough
		case ultipa.PropertyType_LIST:
			listDataRows, err := parseAttrList(oAttr)
			if err != nil {
				return nil, err
			}
			for _, row := range listDataRows {
				attr.Rows = append(attr.Rows, row)
			}
		case ultipa.PropertyType_MAP:
			mapDataRows, err := parseAttrMap(oAttr)
			if err != nil {
				return nil, err
			}
			for _, row := range mapDataRows {
				attr.Rows = append(attr.Rows, row)
			}
		default:
			for _, v := range oAttr.Values {
				value, err := utils.ConvertBytesToInterface(v, attr.PropertyType, nil)
				if err != nil {
					return nil, err
				}
				attr.Rows = append(attr.Rows, value)
			}
		}
	}
	return attr, nil
}

// parseAttrList parse the Attr that PropertyType is ultipa.PropertyType_LIST
func parseAttrList(oAttr *ultipa.Attr) ([]*structs.AttrListData, error) {
	var listDataRows []*structs.AttrListData
	for _, v := range oAttr.Values {
		oListData := &ultipa.AttrListData{}
		listData := structs.NewAttrListData()
		err := proto.Unmarshal(v, oListData)
		if err != nil {
			return nil, err
		}
		listData.ResultType = oListData.Type
		switch oListData.Type {
		case ultipa.ResultType_RESULT_TYPE_ATTR:
			for _, subOAttr := range oListData.Attrs {
				subAttr, err := parseAttr(subOAttr, "")
				if err != nil {
					return nil, err
				}
				listData.Attrs = append(listData.Attrs, subAttr)
			}
		case ultipa.ResultType_RESULT_TYPE_NODE:
			nodes, _, err := NodeTableToNodes(oListData.Nodes, "")
			if err != nil {
				return nil, err
			}
			listData.Nodes = append(listData.Nodes, nodes...)

		case ultipa.ResultType_RESULT_TYPE_EDGE:
			edges, _, err := EdgeTableToEdges(oListData.Edges, "")
			if err != nil {
				return nil, err
			}
			listData.Edges = append(listData.Edges, edges...)

		case ultipa.ResultType_RESULT_TYPE_PATH:
			paths, err := parsePaths(oListData.Paths, "")
			if err != nil {
				return nil, err
			}
			listData.Paths = append(listData.Paths, paths...)
		}
		listDataRows = append(listDataRows, listData)
	}
	return listDataRows, nil
}

// parseAttrMap parse the Attr that PropertyType is ultipa.PropertyType_MAP
func parseAttrMap(oAttr *ultipa.Attr) ([]*structs.AttrMapData, error) {
	var mapDataRows []*structs.AttrMapData
	for _, v := range oAttr.Values {
		oMapData := &ultipa.AttrMapData{}
		mapData := structs.NewAttrMapData()
		err := proto.Unmarshal(v, oMapData)
		if err != nil {
			return nil, err
		}

		key, err := parseAttr(oMapData.Key, "")
		if err != nil {
			return nil, err
		}
		value, err := parseAttr(oMapData.Value, "")
		if err != nil {
			return nil, err
		}
		mapData.Key = key
		mapData.Value = value
		mapDataRows = append(mapDataRows, mapData)
	}
	return mapDataRows, nil
}

//AsGraphs the types will be tables and alias is nodeSchema and edgeSchema
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
		//0:id, 1: name, 2: totalNodes ,3:totalEdges ,4:description ,5:status
		values := row.GetValues()
		graph := structs.Graph{}
		graph.ID = string(values[0])
		graph.Name = string(values[1])
		graph.TotalNodes, _ = utils.Str2Uint64(utils.AsString(values[2]))
		graph.TotalEdges, _ = utils.Str2Uint64(utils.AsString(values[3]))
		graph.Description = string(values[4])
		graph.Status = string(values[5])

		graphs = append(graphs, &graph)
	}

	return graphs, err
}

//AsSchemas the types will be tables and alias is nodeSchema and edgeSchema
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
		TotalIndex = 3
	}

	for _, row := range table.TableRows {
		//0:name, 1: description, 2: json(properties),3:totalNodes, 4:totalEdges
		values := row.GetValues()
		schema := structs.NewSchema(string(values[0]))
		schema.Desc = string(values[1])
		schema.Type = Type
		propertyJson := values[2]
		schema.Total, _ = strconv.Atoi(utils.AsString(values[TotalIndex]))

		schema.DBType, err = structs.GetDBTypeByString(schema.Type)

		if err != nil {
			return nil, err
		}

		var props []*struct {
			Name        string
			Type        string
			Description string
			Lte         string
		}

		err = json.Unmarshal(propertyJson, &props)

		if err != nil {
			log.Fatalln(err)
		}

		for _, prop := range props {
			lte, err := strconv.ParseBool(prop.Lte)
			if err != nil {
				log.Fatalln(err)
			}
			p := structs.Property{
				Name:   prop.Name,
				Desc:   prop.Description,
				Lte:    lte,
				Schema: schema.Name,
			}
			p.SetTypeByString(prop.Type)
			schema.Properties = append(schema.Properties, &p)
		}

		schemas = append(schemas, schema)
	}

	return schemas, err
}

//AsProperties the types will be tables and alias is nodeProperty and edgeProperty
func (di *DataItem) AsProperties() (properties []*structs.Property, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return properties, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " should be a table as pre-condition")
	}

	table := di.Data.(*ultipa.Table)

	if table.TableName != RESP_NODE_PROPERTY_KEY && table.TableName != RESP_EDGE_PROPERTY_KEY {
		return nil, errors.New("DataItem " + di.Alias + " is not a Property list")
	}

	for _, row := range table.TableRows {
		//0:name, 1: type, 2: lte, 3: schema, 4: description
		values := row.GetValues()

		lte, err := strconv.ParseBool(string(values[2]))
		if err != nil {
			log.Fatalln(err)
		}
		p := structs.Property{
			Name:   string(values[0]),
			Desc:   string(values[4]),
			Lte:    lte,
			Schema: string(values[3]),
		}
		p.SetTypeByString(string(values[1]))
		properties = append(properties, &p)

	}

	return properties, err
}

//AsIndexes the types will be tables and alias is nodeIndex and edgeIndex
func (di *DataItem) AsIndexes() (indexes []*structs.Index, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return indexes, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " should be a table as pre-condition")
	}

	table := di.Data.(*ultipa.Table)

	if table.TableName != RESP_NODE_INDEX_KEY && table.TableName != RESP_EDGE_INDEX_KEY {
		return nil, errors.New("DataItem " + di.Alias + " is not a Index list")
	}

	for _, row := range table.TableRows {
		//0:name, 1: properties, 2: schema, 3: status
		values := row.GetValues()

		i := structs.Index{
			Name:       string(values[0]),
			Properties: string(values[1]),
			Schema:     string(values[2]),
			Status:     string(values[3]),
		}
		indexes = append(indexes, &i)

	}

	return indexes, err
}

//AsFullText the types will be tables and alias is node fulltext Index and edge fulltext Index
func (di *DataItem) AsFullText() (fullTextIndexes []*structs.Index, err error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return fullTextIndexes, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_TABLE {
		return nil, errors.New("DataItem " + di.Alias + " should be a table as pre-condition")
	}

	table := di.Data.(*ultipa.Table)

	if table.TableName != RESP_NODE_FULLTEXT_KEY && table.TableName != RESP_EDGE_FULLTEXT_KEY {
		return nil, errors.New("DataItem " + di.Alias + " is not a Fulltext Index list")
	}

	for _, row := range table.TableRows {
		//0:name, 1: properties, 2: schema, 3: status
		values := row.GetValues()

		i := structs.Index{
			Name:       string(values[0]),
			Properties: string(values[1]),
			Schema:     string(values[2]),
			Status:     string(values[3]),
		}
		fullTextIndexes = append(fullTextIndexes, &i)

	}

	return fullTextIndexes, err
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

func (di *DataItem) AsAny() (interface{}, error) {

	switch di.Type {
	case ultipa.ResultType_RESULT_TYPE_ATTR:
		return di.AsAttr()
	//case ultipa.ResultType_RESULT_TYPE_ARRAY:
	//	return di.AsArray()
	case ultipa.ResultType_RESULT_TYPE_EDGE:
		edges, _, err := di.AsEdges()
		return edges, err
	case ultipa.ResultType_RESULT_TYPE_NODE:
		nodes, _, err := di.AsNodes()
		return nodes, err
	case ultipa.ResultType_RESULT_TYPE_TABLE:
		return di.AsTable()
	default:
		return di.Data, nil
	}

}
