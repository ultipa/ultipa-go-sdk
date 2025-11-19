package structs

type GetUserSetting struct {
	UserName string
	Type     string //optional(可选参数)
}
type SetUserSetting struct {
	UserName string
	Type     string //optional(可选参数)
	Data     string //optional(可选参数)
}
