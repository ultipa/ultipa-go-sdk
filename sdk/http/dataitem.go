package http

import (
	"encoding/json"
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"google.golang.org/protobuf/proto"
	"log"
	"strconv"
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
		schema.DBType = ultipa.DBType_DBNODE
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
		schema.DBType = ultipa.DBType_DBEDGE
		schemas[schema.Name] = schema
		for _, header := range oSchema.Properties {
			schema.Properties = append(schema.Properties, &structs.Property{Name: header.PropertyName, Type: header.PropertyType, SubTypes: header.SubTypes})
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

	if di.Type != ultipa.ResultType_RESULT_TYPE_NODE && di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, schemas, errors.New(fmt.Sprintf("dataItem %s is not either Node type or LIST Node type", di.Alias))
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

	if di.Type != ultipa.ResultType_RESULT_TYPE_EDGE && di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, schemas, errors.New(fmt.Sprintf("dataItem %s is not either Edge type or LIST Edge type", di.Alias))
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

	if di.Type != ultipa.ResultType_RESULT_TYPE_PATH && di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New(fmt.Sprintf("dataItem %s is not either Path type or LIST Path type", di.Alias))
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

func (di *DataItem) AsAttr() (*structs.Attr, error) {

	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return nil, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Attribute list")
	}

	attrAlias := di.Data.(*ultipa.AttrAlias)
	oAttr := attrAlias.Attr

	midAttr, err := parseAttr(oAttr, attrAlias.Alias)
	if err != nil {
		return nil, err
	}
	switch midAttr.PropertyType {
	case ultipa.PropertyType_LIST:
		return midAttr.ListAttrAsAttr()

	case ultipa.PropertyType_SET:
		return nil, errors.New(fmt.Sprintf("DataItem %v is not either Type Attr or LIST Attr, but SET, not supported yet.", di.Alias))
	case ultipa.PropertyType_MAP:
		return nil, errors.New(fmt.Sprintf("DataItem %v is not either Type Attr or LIST Attr, but MAP, not supported yet.", di.Alias))
	default:
		return midAttr, nil
	}
	return nil, err
}

// AsAttrEdges parse DataItem as Attr with Rows that is List<List<Node>>
func (di *DataItem) AsAttrNodes() (*structs.AttrNodes, error) {
	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return nil, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Attr")
	}

	attrAlias := di.Data.(*ultipa.AttrAlias)
	oAttr := attrAlias.Attr

	midAttr, err := parseAttr(oAttr, attrAlias.Alias)
	if err != nil {
		return nil, err
	}
	return midAttr.ListAttrAsAttrNodes()
}

// AsAttrEdges parse DataItem as Attr with Rows that is List<List<Edge>>
func (di *DataItem) AsAttrEdges() (*structs.AttrEdges, error) {
	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return nil, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Attr")
	}

	attrAlias := di.Data.(*ultipa.AttrAlias)
	oAttr := attrAlias.Attr

	midAttr, err := parseAttr(oAttr, attrAlias.Alias)
	if err != nil {
		return nil, err
	}
	return midAttr.ListAttrAsAttrEdges()
}

// AsAttrPaths parse DataItem as Attr with Rows that is List<List<Path>>
func (di *DataItem) AsAttrPaths() (*structs.AttrPaths, error) {
	if di.Type == ultipa.ResultType_RESULT_TYPE_UNSET {
		return nil, nil
	}

	if di.Type != ultipa.ResultType_RESULT_TYPE_ATTR {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Attr")
	}

	attrAlias := di.Data.(*ultipa.AttrAlias)
	oAttr := attrAlias.Attr

	midAttr, err := parseAttr(oAttr, attrAlias.Alias)
	if err != nil {
		return nil, err
	}
	return midAttr.ListAttrAsAttrPaths()
}

func parseAttr(oAttr *ultipa.Attr, name string) (*structs.Attr, error) {
	attr := structs.NewAttr()
	attr.Name = name
	attr.PropertyType = oAttr.ValueType
	if oAttr == nil || oAttr.Values == nil {
		return attr, nil
	}

	err := handleAttrValues(oAttr, attr)
	if err != nil {
		return attr, err
	}
	return attr, nil
}

func handleAttrValues(oAttr *ultipa.Attr, attr *structs.Attr) error {
	switch oAttr.ValueType {
	case ultipa.PropertyType_SET:
		fallthrough
	case ultipa.PropertyType_LIST:
		err := parseAttrList(oAttr, attr)
		if err != nil {
			return err
		}
	case ultipa.PropertyType_MAP:
		mapDataRows, err := parseAttrMap(oAttr)
		if err != nil {
			return err
		}
		for _, row := range mapDataRows {
			attr.Rows = append(attr.Rows, row)
		}
	default:
		attr.ResultType = ultipa.ResultType_RESULT_TYPE_ATTR
		if oAttr.Values == nil {
			attr.Rows = nil
		} else {
			for _, v := range oAttr.Values {
				value, err := utils.ConvertBytesToInterface(v, attr.PropertyType, nil)
				if err != nil {
					return err
				}
				attr.Rows = append(attr.Rows, value)
			}
		}
	}
	return nil
}

// parseAttrList parse the oAttr that PropertyType is ultipa.PropertyType_LIST, set the parsed value to attr.
func parseAttrList(oAttr *ultipa.Attr, attr *structs.Attr) error {
	var resultType ultipa.ResultType
	for _, v := range oAttr.Values {
		oListData := &ultipa.AttrListData{}
		err := proto.Unmarshal(v, oListData)
		if err != nil {
			return err
		}
		if oListData.IsNull {
			attr.Rows = append(attr.Rows, nil)
			continue
		}
		switch oListData.Type {
		case ultipa.ResultType_RESULT_TYPE_NODE:
			nodes, _, err := NodeTableToNodes(oListData.Nodes, "")
			if err != nil {
				return err
			}
			listData := structs.NewAttrListData()
			listData.ResultType = oListData.Type
			listData.Nodes = append(listData.Nodes, nodes...)
			attr.Rows = append(attr.Rows, listData)
			if ultipa.ResultType_RESULT_TYPE_UNSET == resultType {
				resultType = listData.ResultType
			}
		case ultipa.ResultType_RESULT_TYPE_EDGE:
			edges, _, err := EdgeTableToEdges(oListData.Edges, "")
			if err != nil {
				return err
			}
			listData := structs.NewAttrListData()
			listData.ResultType = oListData.Type
			listData.Edges = append(listData.Edges, edges...)
			attr.Rows = append(attr.Rows, listData)
			if ultipa.ResultType_RESULT_TYPE_UNSET == resultType {
				resultType = listData.ResultType
			}
		case ultipa.ResultType_RESULT_TYPE_PATH:
			paths, err := parsePaths(oListData.Paths, "")
			if err != nil {
				return err
			}
			listData := structs.NewAttrListData()
			listData.ResultType = oListData.Type
			listData.Paths = append(listData.Paths, paths...)
			attr.Rows = append(attr.Rows, listData)
			if ultipa.ResultType_RESULT_TYPE_UNSET == resultType {
				resultType = listData.ResultType
			}
		case ultipa.ResultType_RESULT_TYPE_ATTR:
			if ultipa.ResultType_RESULT_TYPE_UNSET == resultType {
				resultType = ultipa.ResultType_RESULT_TYPE_ATTR
			}
			//not null but len==0, then set an empty slice
			row := structs.Row{}

			for _, subOAttr := range oListData.Attrs {
				subAttr, err := parseAttr(subOAttr, "")
				if err != nil {
					return err
				}
				if subAttr == nil {
					row = append(row, nil)
				} else {
					row = append(row, subAttr.Rows...)
				}
			}
			attr.Rows = append(attr.Rows, row)

		}
	}
	if ultipa.ResultType_RESULT_TYPE_UNSET == resultType {
		resultType = ultipa.ResultType_RESULT_TYPE_ATTR
	}
	attr.ResultType = resultType
	return nil
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

// AsGraphs the types will be tables and alias is nodeSchema and edgeSchema
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

// AsSchemas the types will be tables and alias is nodeSchema and edgeSchema
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

// AsProperties the types will be tables and alias is nodeProperty and edgeProperty
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

		rowValues := map[string][]byte{}
		for idx, header := range table.Headers {
			rowValues[header.PropertyName] = values[idx]
		}

		name := getOrDefault("name", "", rowValues)
		lteStr := getOrDefault("lte", "false", rowValues)
		typeStr := getOrDefault("type", "", rowValues)
		read := getOrDefault("read", "0", rowValues)
		write := getOrDefault("write", "0", rowValues)
		schema := getOrDefault("schema", "0", rowValues)
		desc := getOrDefault("description", "", rowValues)
		lte, err := strconv.ParseBool(lteStr)
		extra := getOrDefault("extra", "", rowValues)
		if err != nil {
			log.Fatalln(err)
		}
		p := structs.Property{
			Name:   name,
			Desc:   desc,
			Lte:    lte,
			Read:   "1" == read,
			Write:  "1" == write,
			Schema: schema,
			Extra:  extra,
		}
		p.SetTypeByString(typeStr)
		properties = append(properties, &p)

	}

	return properties, err
}

func getOrDefault(name string, defaultValue string, container map[string][]byte) string {
	bytes, ok := container[name]
	if ok {
		return string(bytes)
	}
	return defaultValue
}

// AsIndexes the types will be tables and alias is nodeIndex and edgeIndex
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

// AsFullText the types will be tables and alias is node fulltext Index and edge fulltext Index
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
