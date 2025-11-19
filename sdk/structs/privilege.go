package structs

type PrivilegeType string
type PrivilegeTypes []PrivilegeType
type PrivilegeTargetType string

const (
	Read  PrivilegeType = "READ"
	Write PrivilegeType = "WRITE"
	Deny  PrivilegeType = "DENY"
)

const (
	PrivilegeToUser   PrivilegeTargetType = "user"
	PrivilegeToPolicy PrivilegeTargetType = "policy"
)

type Privilege struct {
	GraphPrivileges  []string
	SystemPrivileges []string
}
