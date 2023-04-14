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
	if strings.Contains(name, "\"") {
		return true
	}
	//name contains multi-byte character
	matcher := regexp.MustCompile("[^\\x00-\\xff]+")
	return matcher.Match([]byte(name))
}

//IsNeedToEscapeSchemaNameForProperty check that whether schema name should be escaped by `` when creating property,
//true - need to escaped, false - no need
func IsNeedToEscapeSchemaNameForProperty(name string) bool {
	return IsNeedToEscapeName(name) || strings.HasPrefix(name, ".") || IsBeginWithDigital(name)
}

func IsBeginWithDigital(name string) bool {
	matcher := regexp.MustCompile("^\\d")
	return matcher.Match([]byte(name))
}
