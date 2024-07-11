package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) ShowPrivilege(config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.Uql("show().privilege()", config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) GrantPropertyPrivilege(privilegeTypes structs.PrivilegeTypes, graph, schema, property string, targetType structs.PrivilegeTargetType, privilegeName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`grant().privilege(%s).on(
  "%s", 
  @%s, 
  "%s"
).%s("%s")`, utils.ToJSONString(privilegeTypes), graph, schema, property, targetType, privilegeName)

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) GrantSystemPrivilege(systemPrivilegeTypes string, targetType structs.PrivilegeTargetType, privilegeName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`grant().system().privilege(["%s"]).%s("%s")`, systemPrivilegeTypes, targetType, privilegeName)

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) RevokePropertyPrivilege(privilegeTypes structs.PrivilegeType, graph, schema, property string, targetType structs.PrivilegeTargetType, privilegeName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`revoke().privilege(["%s"]).on(
  "%s", 
  @%s, 
  "%s"
).%s("%s")`, privilegeTypes, graph, schema, property, targetType, privilegeName)

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) RevokeSystemPrivilege(systemPrivilegeTypes string, targetType structs.PrivilegeTargetType, privilegeName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`revoke().system().privilege(["%s"]).%s("%s")`, systemPrivilegeTypes, targetType, privilegeName)

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
