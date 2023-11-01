package structs

import (
	"encoding/json"
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"regexp"
	"strconv"
	"strings"
)

type Property struct {
	Name     string
	Desc     string
	Lte      bool
	Read     bool
	Write    bool
	Schema   string
	Type     ultipa.PropertyType
	SubTypes []ultipa.PropertyType
	// extra info for property in json format, e.g. DecimalExtra: precision and scale for decimal type.
	Extra string
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
	"_id":        PropertyType_ID,
	"_uuid":      PropertyType_UUID,
	"_from":      PropertyType_FROM,
	"_to":        PropertyType_TO,
	"_from_uuid": PropertyType_FROM_UUID,
	"_to_uuid":   PropertyType_TO_UUID,
	"_ignore":    PropertyType_IGNORE,
	"string":     ultipa.PropertyType_STRING,
	"int32":      ultipa.PropertyType_INT32,
	"int64":      ultipa.PropertyType_INT64,
	"uint32":     ultipa.PropertyType_UINT32,
	"uint64":     ultipa.PropertyType_UINT64,
	"float":      ultipa.PropertyType_FLOAT,
	"double":     ultipa.PropertyType_DOUBLE,
	"datetime":   ultipa.PropertyType_DATETIME,
	"timestamp":  ultipa.PropertyType_TIMESTAMP,
	"text":       ultipa.PropertyType_TEXT,
	"blob":       ultipa.PropertyType_BLOB,
	"point":      ultipa.PropertyType_POINT,
	"decimal":    ultipa.PropertyType_DECIMAL,
	"NULL_":      ultipa.PropertyType_NULL_,
}

var PropertyReverseMap = map[ultipa.PropertyType]string{
	PropertyType_ID:               "_id",
	PropertyType_UUID:             "_uuid",
	PropertyType_FROM:             "_from",
	PropertyType_TO:               "_to",
	PropertyType_FROM_UUID:        "_from_uuid",
	PropertyType_TO_UUID:          "_to_uuid",
	PropertyType_IGNORE:           "_ignore",
	ultipa.PropertyType_STRING:    "string",
	ultipa.PropertyType_INT32:     "int32",
	ultipa.PropertyType_INT64:     "int64",
	ultipa.PropertyType_UINT32:    "uint32",
	ultipa.PropertyType_UINT64:    "uint64",
	ultipa.PropertyType_FLOAT:     "float",
	ultipa.PropertyType_DOUBLE:    "double",
	ultipa.PropertyType_DATETIME:  "datetime",
	ultipa.PropertyType_TIMESTAMP: "timestamp",
	ultipa.PropertyType_TEXT:      "text",
	ultipa.PropertyType_BLOB:      "blob",
	ultipa.PropertyType_POINT:     "point",
	ultipa.PropertyType_DECIMAL:   "decimal",
	ultipa.PropertyType_NULL_:     "NULL_",
}

type DecimalExtra struct {
	Precision int `json:"precision"`
	Scale     int `json:"scale"`
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
	if strings.HasSuffix(s, "[]") {
		p.Type = ultipa.PropertyType_LIST
		p.SubTypes = append(p.SubTypes, GetPropertyTypeByString(strings.TrimSuffix(s, "[]")))
		return
	}
	reg := regexp.MustCompile(`set\(([^)]+)\)`)
	if reg.MatchString(strings.ReplaceAll(s, " ", "")) {
		matches := reg.FindStringSubmatch(strings.ReplaceAll(s, " ", ""))
		p.Type = ultipa.PropertyType_SET
		p.SubTypes = append(p.SubTypes, GetPropertyTypeByString(matches[1]))
		return
	}
	re := regexp.MustCompile(`^decimal\((\d+),(\d+)\)$`)
	if re.MatchString(strings.ReplaceAll(s, " ", "")) {
		p.Type = ultipa.PropertyType_DECIMAL
		matches := re.FindStringSubmatch(strings.ReplaceAll(s, " ", ""))
		precision, err := strconv.Atoi(matches[1])
		scale, err := strconv.Atoi(matches[2])
		if err != nil {
			return
		}
		extraData := DecimalExtra{
			Precision: precision,
			Scale:     scale,
		}
		extraJson, err := json.Marshal(extraData)
		if err != nil {
			return
		}
		p.Extra = string(extraJson)
		return
	}
	p.Type = GetPropertyTypeByString(s)
}

func (p *Property) GetStringType() (string, error) {
	if p.Type == ultipa.PropertyType_LIST {
		if len(p.SubTypes) == 0 {
			return "", errors.New(fmt.Sprintf("Property [%s] is List but not specified subTypes", p.Name))
		}
		return GetStringByPropertyType(p.SubTypes[0]) + "[]", nil
	}
	if p.Type == ultipa.PropertyType_SET {
		if len(p.SubTypes) == 0 {
			return "", errors.New(fmt.Sprintf("Property [%s] is Set but not specified subTypes", p.Name))
		}
		return fmt.Sprintf("set(%s)", GetStringByPropertyType(p.SubTypes[0])), nil
	}
	if p.Type == ultipa.PropertyType_DECIMAL {
		var extraData DecimalExtra
		err := json.Unmarshal([]byte(p.Extra), &extraData)
		if err != nil {
			return GetStringByPropertyType(p.Type), nil
		}
		extraString := fmt.Sprintf("(%d,%d)", extraData.Precision, extraData.Scale)
		return GetStringByPropertyType(p.Type) + extraString, nil
	}
	return GetStringByPropertyType(p.Type), nil
}

func GetPropertyTypeByString(s string) ultipa.PropertyType {
	return PropertyMap[s]
}

func GetStringByPropertyType(t ultipa.PropertyType) string {
	return PropertyReverseMap[t]
}
