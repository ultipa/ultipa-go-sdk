package utils

import "regexp"

// check the name is validated，for PropertyName, GraphName, SchemaName

func CheckGraphName(name string) bool {
	return CheckCustomerName(name)
}

func CheckPropertyName(name string) bool {
	return CheckCustomerName(name)
}

func CheckSchemaName(name string) bool {
	return CheckCustomerName(name)
}

//var CUSTOM_NAME_FORMAT = "Name Should Match REGEXP : ([a-zA-Z][a-zA-Z0-9_]+)|(_id|_uuid|_from|_to|_from_uuid|_to_uuid)"
func CheckCustomerName(name string) bool {
	matcher := regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9_]+)|(_id|_uuid|_from|_to|_from_uuid|_to_uuid)$`)
	return matcher.Match([]byte(name))
}

//CheckCustomerNonIdName check that non-id type property name should match REGEXP: ^([a-zA-Z][a-zA-Z0-9_]+)$
func CheckCustomerNonIdName(name string) bool {
	matcher := regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9_]+)$`)
	return matcher.Match([]byte(name))
}