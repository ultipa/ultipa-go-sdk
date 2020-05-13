package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
)
func FormatNodeAliases(nodes []*ultipa.NodeAlias) *NodeAliases {
	var arrs NodeAliases
	for _, one := range nodes{
		arrs = append(arrs, FormatNodeAlias(one))
	}
	return &arrs
}
func FormatEdgeAliases(edges []*ultipa.EdgeAlias) *EdgeAliases {
	var arrs EdgeAliases
	for _, one := range edges{
		arrs = append(arrs, FormatEdgeAlias(one))
	}
	return &arrs
}
func FormatNodeAlias(node *ultipa.NodeAlias) *NodeAlias {
	newNode := NodeAlias{}
	newNode.Alias = node.Alias
	newNode.NodeTable = FormatNodeTable(node.GetNodeTable())
	return &newNode
}
func FormatEdgeAlias(edge *ultipa.EdgeAlias) *EdgeAlias {
	newEdge := EdgeAlias{}
	newEdge.Alias = edge.Alias
	newEdge.EdgeTable = FormatEdgeTable(edge.GetEdgeTable())
	return &newEdge
}
func FormatNodeTable(nodeTable *ultipa.NodeTable) *NodeTable {
	newTable := NodeTable{}
	var headers []string
	var types []PropertyType
	if nodeTable.Headers != nil {
		for _, v := range nodeTable.Headers {
			headers = append(headers, v.GetPropertyName())
			types = append(types, v.GetPropertyType())
		}
	}
	newTable.Headers = &headers
	var rows []*NodeRow
	if nodeTable.NodeRows != nil {
		for _, v := range nodeTable.GetNodeRows() {
			_node := NodeRow{}
			_node.ID = v.GetId()
			_vs := _formatValues(v.GetValues(), types, headers)
			_node.Values = _vs
			rows = append(rows, &_node)
		}
	}
	newTable.NodeRows = rows
	return &newTable
}
func _formatValues(values [][]byte, types []PropertyType, headers []string) *map[string]interface{}{
	_vs := map[string]interface{}{}
	_missHeaders := false
	if len(values) > len(types) || len(values) > len(headers) {
		_missHeaders = true
		fmt.Println("‼️ BUG 服务器没有返回header")
	}
	for _index, vv := range values {
		vvType := PROPERTY_TYPE_STRING // 服务端有bug，所以，硬修复下 0513
		var key string
		if _missHeaders == false {
			vvType = types[_index]
			key = headers[_index]
		} else {
			key = fmt.Sprintf("Unknown %v", _index)
		}
		change := deserialize(vv, vvType)
		_vs[key] = change
	}
	return &_vs
}
func FormatEdgeTable(edgeTable *ultipa.EdgeTable) *EdgeTable {
	newTable := EdgeTable{}
	var headers []string
	var types []PropertyType
	if edgeTable.Headers != nil {
		for _, v := range edgeTable.Headers {
			headers = append(headers, v.GetPropertyName())
			types = append(types, v.GetPropertyType())
		}
	}
	newTable.Headers = &headers
	var rows []*EdgeRow
	if edgeTable.EdgeRows != nil {
		for _, v := range edgeTable.GetEdgeRows() {
			_edge := EdgeRow{}
			_edge.From = v.GetFromId()
			_edge.ID = v.GetId()
			_edge.To = v.GetToId()
			_vs := _formatValues(v.GetValues(), types, headers)
			_edge.Values = _vs
			rows = append(rows, &_edge)
		}
	}
	newTable.EdgeRows = rows

	return &newTable
}
func FormatPaths(paths []*ultipa.Path) *Paths {
	var ps Paths
	for _,path := range paths{
		newPath := Path{
			NodeTable: FormatNodeTable(path.NodeTable),
			EdgeTable: FormatEdgeTable(path.EdgeTable),
		}
		ps = append(ps, &newPath)
	}
	return &ps
}
func FormatAttrAlias(attrAlias *ultipa.AttrAlias) *AttrAlias  {
	newAttrAlias := AttrAlias{}
	newAttrAlias.Alias = attrAlias.GetAlias()
	if attrAlias.GetValues() != nil {
		var newValues []interface{}
		for _, v := range attrAlias.GetValues(){
			v1 := deserialize(v, attrAlias.GetPropertyType())
			newValues = append(newValues, v1)
		}
		newAttrAlias.Values = newValues
	}
	return &newAttrAlias
}
func FormatAttrs(attrs []*ultipa.AttrAlias) *Attrs  {
	var newAttrs Attrs
	if attrs != nil {
		for _, attr := range attrs {
			newAttrs = append(newAttrs, FormatAttrAlias(attr) )
		}
	}
	return &newAttrs
}
func FormatTables(tables []*ultipa.Table) *Tables  {
	if tables == nil {
		return nil
	}
	var newTables Tables
	for _, table := range tables {
		tb := Table{
			TableName: table.TableName,
			Headers:   table.Headers,
		}
		var newRows TableRows
		tableRows := table.GetTableRows()
		if tableRows != nil {
			for _, row := range tableRows {
				var _row []interface{}
				for _, v := range row.GetValues(){
					_v := deserialize(v, PROPERTY_TYPE_STRING)
					_row = append(_row, _v)
				}
				newRows = append(newRows, &_row)
			}
		}
		tb.TableRows = &newRows
		newTables = append(newTables, &tb)
	}
	return &newTables
}
func FormatKeyValues(values []*ultipa.Value)  *map[string]string{
	if values == nil {
		return nil
	}
	_values := map[string]string{}
	for _, ev := range values {
		_values[ev.GetKey()] = ev.GetValue()
	}
	return &_values
}
func _bytesToRead(bs []byte, out interface{}) {
	//bs[0] = (byte)(int(bs[0]) ^ 0x80)
	buff := bytes.NewBuffer(bs)
	binary.Read(buff, binary.BigEndian, out)
}
func deserialize(bytes []byte, propertyType PropertyType) interface{} {
	switch propertyType {
	case PROPERTY_TYPE_STRING:
		return string(bytes)
	case PROPERTY_TYPE_INT32:
		var num int32
		_bytesToRead(bytes, &num)
		return num
	case PROPERTY_TYPE_INT64:
		var num int64
		_bytesToRead(bytes, &num)
		return num
	case PROPERTY_TYPE_UINT32:
		var num uint32
		_bytesToRead(bytes, &num)
		return num
	case PROPERTY_TYPE_UINT64:
		var num uint64
		_bytesToRead(bytes, &num)
		return num
	case PROPERTY_TYPE_FLOAT:
		var num float32
		_bytesToRead(bytes, &num)
		return num
	case PROPERTY_TYPE_DOUBLE:
		var num float64
		_bytesToRead(bytes, &num)
		return num
	}
	return "Unknown"
}

//
//func FormatEdges(edges []*ultipa.Edge) []*Edge {
//	var newEdges []*Edge
//	for _, ev := range edges {
//		newEdges = append(newEdges, FormatEdge(ev))
//	}
//	return newEdges
//}

//func FormatEdge(edge *ultipa.Edge) *Edge {
//	newEdge := Edge{}
//	newEdge.ID = edge.GetId()
//	newEdge.From = edge.GetFromId()
//	newEdge.To = edge.GetToId()
//	newEdge.Values = FormatValues(edge.GetValues())
//	return &newEdge
//}
func FormatStatus(status *ultipa.Status, err error) *Status {
	if err != nil {
		return &Status{
			Code:    ErrorCode_FAILED,
			Message: fmt.Sprint(err),
		}
	}
	clusterInfo := ClusterInfo{}
	newStatus := Status{
		Code:        ErrorCode_SUCCESS,
		Message:     "",
		ClusterInfo: &clusterInfo,
	}
	newStatus.Code = status.GetErrorCode()
	if status.GetErrorCode() != ultipa.ErrorCode_SUCCESS {
		newStatus.Message = status.GetMsg()
	}
	_clusterInfo := status.GetClusterInfo()
	if _clusterInfo != nil {
		clusterInfo.Redirect = _clusterInfo.Redirect
		clusterInfo.RaftPeers = strings.Split(_clusterInfo.RaftPeers, ",")
	}
	return &newStatus
}

//
//func TableToValues(table *Table) *map[string][]string {
//	res := map[string][]string{}
//	for index, key := range table.Headers {
//		res[key] = table.TableRows[index]
//	}
//	return &res
//}
//
//func TableToArray(table *Table) *[]*map[string]string {
//	var res []*map[string]string
//	for _, rows := range table.TableRows {
//		item := map[string]string{}
//		for index, row := range rows {
//			item[table.Headers[index]] = row
//		}
//		res = append(res, &item)
//	}
//	return &res
//}
