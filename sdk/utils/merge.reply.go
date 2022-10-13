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

//func MergeUQLItems(s1 interface{}, s2 interface{}, fieldName string, tableName string, rowName string) error {
//	var err error
//
//
//	s1v := reflect.ValueOf(s1).Elem().FieldByName(fieldName)
//	s2v := reflect.ValueOf(s2).Elem().FieldByName(fieldName)
//
//	if s1v.IsNil() {
//		s1v.Set(s2v)
//	}
//
//
//	for i := 0; i < s2v.Len(); i++ {
//		alias := ""
//
//		if fieldName == "Tables" {
//			alias = s2v.Field(i).Elem().FieldByName("TableName").Interface().(string)
//		} else {
//			alias = s2v.Field(i).Elem().FieldByName("Alias").Interface().(string)
//		}
//
//		if alias == "" {
//			continue
//		}
//
//		// find same alias and set values
//		found := false
//		for ii :=0; ii < s1v.Len(); ii++ {
//			alias2 := ""
//			if fieldName == "Tables" {
//				alias2 = s2v.Field(i).Elem().FieldByName("TableName").Interface().(string)
//			} else {
//				alias2 = s2v.Field(i).Elem().FieldByName("Alias").Interface().(string)
//			}
//
//			if alias == alias2 {
//				v1 := s1v.Field(i).Elem()
//				v2 := s2v.Field(ii).Elem()
//
//				if fieldName == "Tables" {
//					v1.Set(reflect.AppendSlice(v1.FieldByName(rowName), v2.FieldByName(rowName)))
//				} else {
//					v11 := v1.FieldByName(tableName)
//					v22 := v2.FieldByName(tableName)
//					v22.Set(reflect.AppendSlice(v11.FieldByName(rowName), v22.FieldByName(rowName)))
//				}
//
//				found = true
//			}
//		}
//
//		if found == false {
//			s1v.Set(reflect.AppendSlice(s1v, s2v))
//		}
//
//	}
//
//
//
//
//	return err
//}

func MergeUQLReply(reply1 *ultipa.UqlReply, reply2 *ultipa.UqlReply) *ultipa.UqlReply {

	//err := MergeSameStruct(reply1, reply2)

	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//return reply1

	for _, Alias := range reply1.Alias {
		switch Alias.ResultType {
		case ultipa.ResultType_RESULT_TYPE_NODE:
			data1 := Find(reply1.Nodes, func(index int) bool { return reply1.Nodes[index].Alias == Alias.Alias }).(*ultipa.NodeAlias)
			data2 := Find(reply2.Nodes, func(index int) bool { return reply2.Nodes[index].Alias == Alias.Alias }).(*ultipa.NodeAlias)

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

			table1 := Find(reply1.Tables, func(index int) bool { return reply1.Tables[index].TableName == Alias.Alias })
			table2 := Find(reply2.Tables, func(index int) bool { return reply2.Tables[index].TableName == Alias.Alias })

			if table2 == nil {
				continue
			}

			if table1 == nil && table2 != nil {
				reply1.Tables = append(reply1.Tables, table2.(*ultipa.Table))
				continue
			}
			data1 := table1.(*ultipa.Table)
			data2 := table2.(*ultipa.Table)

			data1.TableRows = append(data1.TableRows, data2.TableRows...)

		case ultipa.ResultType_RESULT_TYPE_PATH:
			if(reply2.Paths == nil ) {
				continue
			}
			data1 := Find(reply1.Paths, func(index int) bool { return reply1.Paths[index].Alias == Alias.Alias }).(*ultipa.PathAlias)
			data2 := Find(reply2.Paths, func(index int) bool { return reply2.Paths[index].Alias == Alias.Alias }).(*ultipa.PathAlias)
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
