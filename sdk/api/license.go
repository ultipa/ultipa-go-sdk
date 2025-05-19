package api

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowLicense(requestConfig *configuration.RequestConfig) (license *structs.License, err error) {
	uql := fmt.Sprintf("license.dump()")
	return api.license(uql, requestConfig)
}

func (api *UltipaAPI) license(uql string, requestConfig *configuration.RequestConfig) (license *structs.License, err error) {
	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	table, err := resp.Alias(http.RESP_LICENSE_KEY).AsTable()
	if err != nil {
		return nil, err
	}

	values := table.ToKV()[0]
	license = &structs.License{
		LicenseUUID: values.Get("license_uuid").(string),
		Company:     values.Get("company").(string),
		Department:  values.Get("department").(string),
		ExpiredDate: values.Get("expired_date").(string),
	}

	return license, nil
}
