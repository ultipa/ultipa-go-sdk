package structs

type GetUserSetting struct {
	UserName string
	Type     string //optional
}
type SetUserSetting struct {
	UserName string
	Type     string //optional
	Data     string //optional
}
