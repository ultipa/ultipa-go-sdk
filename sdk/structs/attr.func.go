package structs

import (
	"errors"
	"fmt"
	ultipa "ultipa-go-sdk/rpc"
)

//ListAttrAsNodes returns node list, if PropertyType of attr is LIST and inner result type is Node
func (attr *Attr) ListAttrAsNodes() ([]*Node, map[string]*Schema, error) {
	if ultipa.PropertyType_LIST != attr.PropertyType {
		return nil, nil, errors.New(fmt.Sprintf("value of this %v is not a LIST type", attr.Name))
	}
	var result []*Node
	if attr.Rows == nil {
		return result, nil, nil
	}
	if len(attr.Rows) == 0 {
		result = []*Node{}
		return result, nil, nil
	}

	detectedInnerResultType := attr.detectListAttrInnerResultType()
	if ultipa.ResultType_RESULT_TYPE_NODE != detectedInnerResultType {
		return nil, nil, errors.New(fmt.Sprintf("value of this %v is not a LIST Node type", attr.Name))
	}

	for _, row := range attr.Rows {
		attrListData := row.(*AttrListData)
		result = append(result, attrListData.Nodes...)
	}
	return result, GetSchemasOfNodeList(result), nil
}

//ListAttrAsEdges returns edge list, if PropertyType of attr is LIST and inner result type is Edge
func (attr *Attr) ListAttrAsEdges() ([]*Edge, map[string]*Schema, error) {
	if ultipa.PropertyType_LIST != attr.PropertyType {
		return nil, nil, errors.New(fmt.Sprintf("value of this %v is not a LIST type", attr.Name))
	}
	var result []*Edge
	if attr.Rows == nil {
		return result, nil, nil
	}
	if len(attr.Rows) == 0 {
		result = []*Edge{}
		return result, nil, nil
	}

	detectedInnerResultType := attr.detectListAttrInnerResultType()
	if ultipa.ResultType_RESULT_TYPE_EDGE != detectedInnerResultType {
		return nil, nil, errors.New(fmt.Sprintf("value of this %v is not a LIST Edge type", attr.Name))
	}

	for _, row := range attr.Rows {
		attrListData := row.(*AttrListData)
		result = append(result, attrListData.Edges...)
	}
	return result, GetSchemasOfEdgeList(result), nil
}

//ListAttrAsPaths returns path list, if PropertyType of attr is LIST and inner result type is Path
func (attr *Attr) ListAttrAsPaths() ([]*Path, error) {
	if ultipa.PropertyType_LIST != attr.PropertyType {
		return nil, errors.New(fmt.Sprintf("value of this %v is not a LIST type", attr.Name))
	}
	var result []*Path
	if attr.Rows == nil {
		return result, nil
	}
	if len(attr.Rows) == 0 {
		result = []*Path{}
		return result, nil
	}

	detectedInnerResultType := attr.detectListAttrInnerResultType()

	if ultipa.ResultType_RESULT_TYPE_PATH != detectedInnerResultType {
		return nil, errors.New(fmt.Sprintf("value of this %v is not a LIST Path type", attr.Name))
	}

	for _, row := range attr.Rows {
		attrListData := row.(*AttrListData)
		result = append(result, attrListData.Paths...)
	}
	return result, nil

}

//ListAttrAsAttr returns an attr, if PropertyType of attr is LIST and inner result type is Attr.
//inner result type is attr, then regarded it as basic type, e.g. string,uint64, float64 etc.
func (attr *Attr) ListAttrAsAttr() (*Attr, error) {
	if ultipa.PropertyType_LIST != attr.PropertyType {
		return nil, errors.New(fmt.Sprintf("value of this %v is not a LIST type", attr.Name))
	}
	result := NewAttr()
	result.Name = attr.Name
	if attr.Rows == nil {
		result.Rows = nil
		return result, nil
	}
	if len(attr.Rows) == 0 {
		return result, nil
	}

	detectedInnerResultType := attr.detectListAttrInnerResultType()
	// if inner result type is attr, then regarded it as basic type, e.g. string,uint64, float64 etc.
	if ultipa.ResultType_RESULT_TYPE_ATTR == detectedInnerResultType {
		for _, row := range attr.Rows {
			var resultRow Row
			attrListData := row.(*AttrListData)
			if attrListData.Attrs == nil {
				resultRow = nil
			} else if len(attrListData.Attrs) == 0 {
				resultRow = Row{}
			} else {
				for _, innerAttr := range attrListData.Attrs {
					resultRow = append(resultRow, innerAttr.Rows...)
					if innerAttr.PropertyType != 0 && result.PropertyType == 0 {
						result.PropertyType = innerAttr.PropertyType
					}
				}
			}
			result.Rows = append(result.Rows, resultRow)
		}
		return result, nil
	}
	return nil, errors.New(fmt.Sprintf("value of this %v is not a LIST Attr type", attr.Name))
}

//@depreacated
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

// detectListAttrInnerResultType detects inner result type of Attr with List property type.
func (attr *Attr) detectListAttrInnerResultType() ultipa.ResultType {
	for _, row := range attr.Rows {
		toDetectTypeAttrListData := row.(*AttrListData)
		if toDetectTypeAttrListData.ResultType != ultipa.ResultType_RESULT_TYPE_UNSET {
			return toDetectTypeAttrListData.ResultType
		}
	}
	return ultipa.ResultType_RESULT_TYPE_UNSET
}
