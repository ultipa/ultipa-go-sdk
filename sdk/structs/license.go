package structs

type License struct {
	LicenseUUID string `json:"license_uuid"`
	Company     string `json:"company"`
	Department  string `json:"department"`
	ExpiredDate string `json:"expired_date"`
}
