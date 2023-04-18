package utils

import (
	"regexp"
	"strings"
)

//CheckGraphName check the name is validated，for PropertyName, GraphName, SchemaName
// deprecated since 4.2.28
func CheckGraphName(name string) bool {
	return CheckCustomerName(name)
}

// CheckPropertyName check property name
// deprecated since 4.2.28
func CheckPropertyName(name string) bool {
	return CheckCustomerName(name)
}

// CheckSchemaName check schema name
// deprecated since 4.2.28
func CheckSchemaName(name string) bool {
	return CheckCustomerName(name)
}

//CheckCustomerName var CUSTOM_NAME_FORMAT = "Name Should Match REGEXP : ([a-zA-Z][a-zA-Z0-9_]+)|(_id|_uuid|_from|_to|_from_uuid|_to_uuid)"
// deprecated since 4.2.28
func CheckCustomerName(name string) bool {
	matcher := regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9_]+)|(_id|_uuid|_from|_to|_from_uuid|_to_uuid)$`)
	return matcher.Match([]byte(name))
}

//CheckCustomerNonIdName check that non-id type property name should match REGEXP: ^([a-zA-Z][a-zA-Z0-9_]+)$
func CheckCustomerNonIdName(name string) bool {
	matcher := regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9_]+)$`)
	return matcher.Match([]byte(name))
}

//IsNeedToEscapeName check that whether name should be escaped by ``,
//true - need to escaped, false - no need
func IsNeedToEscapeName(name string) bool {
	//if strings.Contains(name, "\"") {
	//	return true
	//}
	////name contains multi-byte character
	//matcher := regexp.MustCompile("[^\\x00-\\xff]+")
	//return matcher.Match([]byte(name))
	return IsNeedToEscapeSchemaNameForProperty(name)
}

//IsNeedToEscapeSchemaNameForProperty check that whether schema name should be escaped by `` when creating property,
//true - need to escaped, false - no need
func IsNeedToEscapeSchemaNameForProperty(name string) bool {
	matcher := regexp.MustCompile(`^([a-zA-Z_][a-zA-Z0-9_]+)$`)
	return !matcher.Match([]byte(name))
}

func IsBeginWithDigital(name string) bool {
	matcher := regexp.MustCompile("^\\d")
	return matcher.Match([]byte(name))
}

// StartWithTilde check whether start with ~
func StartWithTilde(name string) bool {
	return strings.HasPrefix(name, "~")
}

var InvalidName = map[string]struct{}{
	"this":          {},
	"prev_n":        {},
	"prev_e":        {},
	"_uuid":         {},
	"_from_uuid":    {},
	"_to_uuid":      {},
	"_id":           {},
	"_from":         {},
	"_to":           {},
	"_graph":        {},
	"_nodeSchema":   {},
	"_edgeSchema":   {},
	"_nodeProperty": {},
	"_edgeProperty": {},
	"_nodeIndex":    {},
	"_edgeIndex":    {},
	"_nodeFulltext": {},
	"_edgeFulltext": {},
	"_statistic":    {},
	"_top":          {},
	"_task":         {},
	"_policy":       {},
	"_user":         {},
	"_privilege":    {},
	"_algoList":     {},
	"_extaList":     {},
}

// IsValidName check whether name is valid or not
func IsValidName(name string) bool {
	if len(name) < 2 || len(name) > 64 {
		return false
	}

	if strings.Contains(name, "`") {
		return false
	}

	if StartWithTilde(name) {
		return false
	}

	if _, ok := InvalidName[name]; ok {
		return false
	}
	return true
}
