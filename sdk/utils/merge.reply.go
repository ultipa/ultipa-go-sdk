package utils

import (
	ultipa "ultipa-go-sdk/rpc"
)

func CheckAliasExist(reply *ultipa.UqlReply, alias string) bool {
	for _, a := range reply.Alias {
		if a.Alias == alias {
			return true
		}
	}

	return false
}

func FindAliasDataInReply(reply *ultipa.UqlReply, alias string) (data interface{}, t ultipa.ResultType) {

	if CheckAliasExist(reply, alias) == false {
		return nil, ultipa.ResultType_RESULT_TYPE_UNSET
	}

	for _, Alias := range reply.Alias {
		switch Alias.ResultType {
		case ultipa.ResultType_RESULT_TYPE_NODE:
			data = Find(reply.Nodes, func(index int) bool { return reply.Nodes[index].Alias == alias })
			t = Alias.ResultType
		case ultipa.ResultType_RESULT_TYPE_EDGE:
			data = Find(reply.Edges, func(index int) bool { return reply.Edges[index].Alias == alias })
			t = Alias.ResultType
		case ultipa.ResultType_RESULT_TYPE_TABLE:
			data = Find(reply.Tables, func(index int) bool { return reply.Tables[index].TableName == alias })
			t = Alias.ResultType
		case ultipa.ResultType_RESULT_TYPE_PATH:
			data = Find(reply.Paths, func(index int) bool { return reply.Paths[index].Alias == alias })
			t = Alias.ResultType
		case ultipa.ResultType_RESULT_TYPE_ARRAY:
			data = Find(reply.Arrays, func(index int) bool { return reply.Arrays[index].Alias == alias })
			t = Alias.ResultType
		case ultipa.ResultType_RESULT_TYPE_ATTR:
			data = Find(reply.Attrs, func(index int) bool { return reply.Attrs[index].Alias == alias })
			t = Alias.ResultType
		default:
			panic("FindAliasDataInReply Not Supported Type")
		}

		if data != nil {
			break
		}
	}

	return data, t

}

func MergeUQLReply(reply1 *ultipa.UqlReply, reply2 *ultipa.UqlReply) *ultipa.UqlReply {

	for _, Alias := range reply1.Alias {
		switch Alias.ResultType {
		case ultipa.ResultType_RESULT_TYPE_NODE:
			data1 := Find(reply1.Nodes, func(index int) bool { return reply1.Nodes[index].Alias == Alias.Alias }).(*ultipa.NodeAlias)
			data2 := Find(reply2.Nodes, func(index int) bool { return reply2.Nodes[index].Alias == Alias.Alias }).(*ultipa.NodeAlias)
			// merge to reply1
			if data2 != nil {
				data1.NodeTable.NodeRows = append(data1.NodeTable.NodeRows, data2.NodeTable.NodeRows...)
			}
		case ultipa.ResultType_RESULT_TYPE_EDGE:
			data1 := Find(reply1.Edges, func(index int) bool { return reply1.Edges[index].Alias == Alias.Alias }).(*ultipa.EdgeAlias)
			data2 := Find(reply2.Edges, func(index int) bool { return reply2.Edges[index].Alias == Alias.Alias }).(*ultipa.EdgeAlias)
			if data2 != nil {
				data1.EdgeTable.EdgeRows = append(data1.EdgeTable.EdgeRows, data2.EdgeTable.EdgeRows...)
			}
		case ultipa.ResultType_RESULT_TYPE_TABLE:
			data1 := Find(reply1.Edges, func(index int) bool { return reply1.Edges[index].Alias == Alias.Alias }).(*ultipa.Table)
			data2 := Find(reply2.Edges, func(index int) bool { return reply2.Edges[index].Alias == Alias.Alias }).(*ultipa.Table)
			if data2 != nil {
				data1.TableRows = append(data1.TableRows, data2.TableRows...)
			}
		case ultipa.ResultType_RESULT_TYPE_PATH:
			data1 := Find(reply1.Edges, func(index int) bool { return reply1.Edges[index].Alias == Alias.Alias }).(*ultipa.PathAlias)
			data2 := Find(reply2.Edges, func(index int) bool { return reply2.Edges[index].Alias == Alias.Alias }).(*ultipa.PathAlias)
			if data2 != nil {
				data1.Paths = append(data1.Paths, data2.Paths...)
			}
		case ultipa.ResultType_RESULT_TYPE_ARRAY:
			data1 := Find(reply1.Arrays, func(index int) bool { return reply1.Arrays[index].Alias == Alias.Alias }).(*ultipa.ArrayAlias)
			data2 := Find(reply2.Arrays, func(index int) bool { return reply2.Arrays[index].Alias == Alias.Alias }).(*ultipa.ArrayAlias)
			if data2 != nil {
				data1.Elements = append(data1.Elements, data2.Elements...)
			}
		case ultipa.ResultType_RESULT_TYPE_ATTR:
			data1 := Find(reply1.Attrs, func(index int) bool { return reply1.Attrs[index].Alias == Alias.Alias }).(*ultipa.AttrAlias)
			data2 := Find(reply2.Attrs, func(index int) bool { return reply2.Attrs[index].Alias == Alias.Alias }).(*ultipa.AttrAlias)
			if data2 != nil {
				data1.Values = append(data1.Values, data2.Values...)
			}
		}
	}

	return reply1
}
