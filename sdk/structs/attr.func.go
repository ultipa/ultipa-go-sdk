package structs

import (
	"errors"
	ultipa "ultipa-go-sdk/rpc"
)

//AsNodes returns node list if PropertyType of Attr is LIST & result type is Node
func (attr *Attr) AsNodes() ([]*Node, error) {
	if ultipa.PropertyType_NULL_ == attr.PropertyType || attr.Rows == nil {
		return nil, nil
	}
	if len(attr.Rows) == 0 {
		return []*Node{}, nil
	}

	if ultipa.PropertyType_LIST == attr.PropertyType {
		attrListData := attr.Rows[0].(*AttrListData)

		if ultipa.ResultType_RESULT_TYPE_NODE != attrListData.ResultType {
			return nil, errors.New("value in list is not *Node type")
		}

		var result []*Node
		for _, row := range attr.Rows {
			attrListData := row.(*AttrListData)
			result = append(result, attrListData.Nodes...)
		}
		return result, nil
	}
	return nil, errors.New("value of this attr is not a []*Node type")
}

//AsEdges returns edge list if PropertyType of Attr is LIST & result type is Edge
func (attr *Attr) AsEdges() ([]*Edge, error) {
	if ultipa.PropertyType_NULL_ == attr.PropertyType || attr.Rows == nil {
		return nil, nil
	}
	if len(attr.Rows) == 0 {
		return []*Edge{}, nil
	}

	if ultipa.PropertyType_LIST == attr.PropertyType {
		attrListData := attr.Rows[0].(*AttrListData)

		if ultipa.ResultType_RESULT_TYPE_EDGE != attrListData.ResultType {
			return nil, errors.New("value in list is not *Edge type")
		}

		var result []*Edge
		for _, row := range attr.Rows {
			attrListData := row.(*AttrListData)
			result = append(result, attrListData.Edges...)
		}
		return result, nil
	}
	return nil, errors.New("value of this attr is not a []*Edge type")
}

//AsPaths returns path list if PropertyType of Attr is LIST & result type is Path
func (attr *Attr) AsPaths() ([]*Path, error) {
	if ultipa.PropertyType_NULL_ == attr.PropertyType || attr.Rows == nil {
		return nil, nil
	}
	if len(attr.Rows) == 0 {
		return []*Path{}, nil
	}

	if ultipa.PropertyType_LIST == attr.PropertyType {
		attrListData := attr.Rows[0].(*AttrListData)

		if ultipa.ResultType_RESULT_TYPE_PATH != attrListData.ResultType {
			return nil, errors.New("value in list is not *Path type")
		}

		var result []*Path
		for _, row := range attr.Rows {
			attrListData := row.(*AttrListData)
			result = append(result, attrListData.Paths...)
		}
		return result, nil
	}
	return nil, errors.New("value of this attr is not a []*Path type")
}

//AsAttr returns attr list if PropertyType of Attr is LIST & result type is Attr
func (attr *Attr) AsAttr() ([]interface{}, error) {
	if ultipa.PropertyType_NULL_ == attr.PropertyType || attr.Rows == nil {
		return nil, nil
	}
	if len(attr.Rows) == 0 {
		return []interface{}{}, nil
	}

	if ultipa.PropertyType_LIST == attr.PropertyType {
		attrListData := attr.Rows[0].(*AttrListData)

		if ultipa.ResultType_RESULT_TYPE_ATTR != attrListData.ResultType {
			return nil, errors.New("value in list is not *Attr type")
		}
		return attr.parseAttrOfAttrListDataToInterface(), nil
	}
	return nil, errors.New("value of this attr is not a []*Attr type")
}

func (attr *Attr) parseAttrOfAttrListDataToInterface() []interface{} {
	var result []interface{}
	switch attr.PropertyType {
	case ultipa.PropertyType_LIST:
		var subAttrs []*Attr

		for _, row := range attr.Rows {
			eachAttrListData := row.(*AttrListData)
			subAttrs = append(subAttrs, eachAttrListData.Attrs...)
		}
		for _, subAttr := range subAttrs {
			subResult := subAttr.parseAttrOfAttrListDataToInterface()
			result = append(result, subResult...)
		}
	case ultipa.PropertyType_SET:
	case ultipa.PropertyType_MAP:
	case ultipa.PropertyType_NULL_:
		result = append(result, nil)
	default:
		//non-collection type, return List<>
		result = append(result, attr.Rows...)
	}
	return result
}
