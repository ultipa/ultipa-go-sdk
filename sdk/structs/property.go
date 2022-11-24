package structs

import ultipa "ultipa-go-sdk/rpc"

type Property struct {
	Name   string
	Desc   string
	Lte    bool
	Schema string
	Type   ultipa.PropertyType
}

const (
	PropertyType_ID        ultipa.PropertyType = -1
	PropertyType_UUID      ultipa.PropertyType = -2
	PropertyType_FROM      ultipa.PropertyType = -3
	PropertyType_TO        ultipa.PropertyType = -4
	PropertyType_FROM_UUID ultipa.PropertyType = -5
	PropertyType_TO_UUID   ultipa.PropertyType = -6
	PropertyType_IGNORE    ultipa.PropertyType = -7
)

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
	"_id":        PropertyType_ID,
	"_uuid":      PropertyType_UUID,
	"_from":      PropertyType_FROM,
	"_to":        PropertyType_TO,
	"_from_uuid": PropertyType_FROM_UUID,
	"_to_uuid":   PropertyType_TO_UUID,
	"_ignore":    PropertyType_IGNORE,
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
	PropertyType_ID:               "_id",
	PropertyType_UUID:             "_uuid",
	PropertyType_FROM:             "_from",
	PropertyType_TO:               "_to",
	PropertyType_FROM_UUID:        "_from_uuid",
	PropertyType_TO_UUID:          "_to_uuid",
	PropertyType_IGNORE:           "_ignore",
}

func (p *Property) IsIDType() bool {

	idTyps := []ultipa.PropertyType{
		PropertyType_ID,
		PropertyType_UUID,
		PropertyType_FROM,
		PropertyType_TO,
		PropertyType_FROM_UUID,
		PropertyType_TO_UUID,
	}

	for _, t := range idTyps {
		if p.Type == t {
			return true
		}
	}

	return false
}

func (p *Property) IsIgnore() bool {
	return p.Type == PropertyType_IGNORE
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
