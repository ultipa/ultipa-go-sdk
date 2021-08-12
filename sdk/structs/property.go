package structs

import ultipa "ultipa-go-sdk/rpc"

type Property struct {
	Name string
	Desc string
	Type ultipa.UltipaPropertyType
}

const (
	UltipaPropertyType_ID       ultipa.UltipaPropertyType = 20
	UltipaPropertyType_UUID     ultipa.UltipaPropertyType = 21
	UltipaPropertyType_FROM     ultipa.UltipaPropertyType = 22
	UltipaPropertyType_FROMUUID ultipa.UltipaPropertyType = 23
	UltipaPropertyType_TO       ultipa.UltipaPropertyType = 24
	UltipaPropertyType_TOUUID   ultipa.UltipaPropertyType = 25
)

var PropertyMap = map[string]ultipa.UltipaPropertyType{
	"string":     ultipa.UltipaPropertyType_STRING,
	"int32":      ultipa.UltipaPropertyType_INT32,
	"int64":      ultipa.UltipaPropertyType_INT64,
	"uint32":     ultipa.UltipaPropertyType_UINT32,
	"uint64":     ultipa.UltipaPropertyType_UINT64,
	"float":      ultipa.UltipaPropertyType_FLOAT,
	"double":     ultipa.UltipaPropertyType_DOUBLE,
	"datetime":   ultipa.UltipaPropertyType_DATETIME,
	"_id":        UltipaPropertyType_ID,
	"_uuid":      UltipaPropertyType_UUID,
	"_from":      UltipaPropertyType_FROM,
	"_to":        UltipaPropertyType_TO,
	"_from_uuid": UltipaPropertyType_FROMUUID,
	"_to_uuid":   UltipaPropertyType_TOUUID,
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
	UltipaPropertyType_ID:              "_id",
	UltipaPropertyType_UUID:            "_uuid",
	UltipaPropertyType_FROM:            "_from",
	UltipaPropertyType_TO:              "_to",
	UltipaPropertyType_FROMUUID:        "_from_uuid",
	UltipaPropertyType_TOUUID:          "_to_uuid",
}

func (p *Property) IsIDType() bool {

	idTyps := []ultipa.UltipaPropertyType{
		UltipaPropertyType_ID,
		UltipaPropertyType_UUID,
		UltipaPropertyType_FROM,
		UltipaPropertyType_TO,
		UltipaPropertyType_FROMUUID,
		UltipaPropertyType_TOUUID,
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

func GetPropertyTypeByString(s string) ultipa.UltipaPropertyType {
	return PropertyMap[s]
}

func GetStringByPropertyType(t ultipa.UltipaPropertyType) string {
	return PropertyReverseMap[t]
}
