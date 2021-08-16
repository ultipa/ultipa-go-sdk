package structs

import ultipa "ultipa-go-sdk/rpc"

type Property struct {
	Name string
	Desc string
	Type ultipa.PropertyType
}

var PropertyMap = map[string]ultipa.PropertyType{
	"string":     ultipa.PropertyType_STRING,
	"int32":      ultipa.PropertyType_INT32,
	"int64":      ultipa.PropertyType_INT64,
	"uint32":     ultipa.PropertyType_UINT32,
	"uint64":     ultipa.PropertyType_UINT64,
	"float":      ultipa.PropertyType_FLOAT,
	"double":     ultipa.PropertyType_DOUBLE,
	"datetime":   ultipa.PropertyType_DATETIME,
	"timestamp":  ultipa.PropertyType_TIMESTAMP,
	"_id":        ultipa.PropertyType_ID,
	"_uuid":      ultipa.PropertyType_UUID,
	"_from":      ultipa.PropertyType_FROM,
	"_to":        ultipa.PropertyType_TO,
	"_from_uuid": ultipa.PropertyType_FROM_UUID,
	"_to_uuid":   ultipa.PropertyType_TO_UUID,
	"_ignore":    ultipa.PropertyType_IGNORE,
}

var PropertyReverseMap = map[ultipa.PropertyType]string{
	ultipa.PropertyType_STRING:    "string",
	ultipa.PropertyType_INT32:     "int32",
	ultipa.PropertyType_INT64:     "int64",
	ultipa.PropertyType_UINT32:    "uint32",
	ultipa.PropertyType_UINT64:    "uint64",
	ultipa.PropertyType_FLOAT:     "float",
	ultipa.PropertyType_DOUBLE:    "double",
	ultipa.PropertyType_DATETIME:  "datetime",
	ultipa.PropertyType_TIMESTAMP: "timestamp",
	ultipa.PropertyType_ID:        "_id",
	ultipa.PropertyType_UUID:      "_uuid",
	ultipa.PropertyType_FROM:      "_from",
	ultipa.PropertyType_TO:        "_to",
	ultipa.PropertyType_FROM_UUID: "_from_uuid",
	ultipa.PropertyType_TO_UUID:   "_to_uuid",
	ultipa.PropertyType_IGNORE:    "_ignore",
}

func (p *Property) IsIDType() bool {

	idTyps := []ultipa.PropertyType{
		ultipa.PropertyType_ID,
		ultipa.PropertyType_UUID,
		ultipa.PropertyType_FROM,
		ultipa.PropertyType_TO,
		ultipa.PropertyType_FROM_UUID,
		ultipa.PropertyType_TO_UUID,
	}

	for _, t := range idTyps {
		if p.Type == t {
			return true
		}
	}

	return false
}

func (p *Property) SetTypeByString(s string) {
	p.Type = GetPropertyTypeByString(s)
}

func (p *Property) GetStringType() string {
	return GetStringByPropertyType(p.Type)
}

func GetPropertyTypeByString(s string) ultipa.PropertyType {
	return PropertyMap[s]
}

func GetStringByPropertyType(t ultipa.PropertyType) string {
	return PropertyReverseMap[t]
}
