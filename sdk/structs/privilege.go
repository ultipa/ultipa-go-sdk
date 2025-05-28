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

type PrivilegeLevel int

const (
	GraphPrivilege PrivilegeLevel = iota
	SystemPrivilege
)

type Privilege struct {
	Name  string
	Level PrivilegeLevel
}
