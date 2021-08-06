package structs

import ultipa "ultipa-go-sdk/rpc"

type Property struct {
	Name string
	Desc string
	Type ultipa.UltipaPropertyType
}

var PropertyMap = map[string]ultipa.UltipaPropertyType{
	"string":   ultipa.UltipaPropertyType_STRING,
	"int32":    ultipa.UltipaPropertyType_INT32,
	"int64":    ultipa.UltipaPropertyType_INT64,
	"uint32":   ultipa.UltipaPropertyType_UINT32,
	"uint64":   ultipa.UltipaPropertyType_UINT64,
	"float":    ultipa.UltipaPropertyType_FLOAT,
	"double":   ultipa.UltipaPropertyType_DOUBLE,
	"datetime": ultipa.UltipaPropertyType_DATETIME,
}

var PropertyReverseMap = map[ultipa.UltipaPropertyType]string{
	ultipa.UltipaPropertyType_STRING:   "string",
	ultipa.UltipaPropertyType_INT32:    "int32",
	ultipa.UltipaPropertyType_INT64:    "int64",
	ultipa.UltipaPropertyType_UINT32:   "uint32",
	ultipa.UltipaPropertyType_UINT64:   "uint64",
	ultipa.UltipaPropertyType_FLOAT:    "float",
	ultipa.UltipaPropertyType_DOUBLE:   "double",
	ultipa.UltipaPropertyType_DATETIME: "datetime",
}

func (p *Property) SetTypeByString(s string) {
	p.Type = GetPropertyTypeByString(s)
}

func (p *Property) GetStringType() string {
	return GetStringByPropertyType(p.Type)
}

func GetPropertyTypeByString(s string) ultipa.UltipaPropertyType {
	return PropertyMap[s]
}

func GetStringByPropertyType(t ultipa.UltipaPropertyType) string {
	return PropertyReverseMap[t]
}