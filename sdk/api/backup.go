package api

import (
	"encoding/json"
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) ShowBacukup(requestConfig *configuration.RequestConfig) (backupinfo []*structs.BackupInfo, err error) {
	uql := fmt.Sprintf("db.backup.show()")

	return api.backup(uql, requestConfig)
}

func (api *UltipaAPI) backup(uql string, requestConfig *configuration.RequestConfig) (backupinfos []*structs.BackupInfo, err error) {
	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	table, err := resp.Alias(http.RESP_BACKUP_KEY).AsTable()
	if err != nil {
		return nil, err
	}

	//printers.PrintTable(table)

	for _, values := range table.ToKV() {
		var backup []structs.ShardInfo
		status := values.Get("status").(string)
		if status == "DONE" {
			bf := values.Get("backup_infos").(string)
			err = json.Unmarshal([]byte(bf), &backup)
			if err != nil {
				api.Logger.Warn(bf + ": backup_infos Unmarshal error" + err.Error())
			}
		}

		var startTime, endTime string
		if values.Get("start_time") != nil {
			startTime = values.Get("start_time").(*utils.UltipaTime).String()
		}

		if values.Get("end_time") != nil {
			endTime = values.Get("end_time").(*utils.UltipaTime).String()
		}

		backupinfo := &structs.BackupInfo{
			BackupName:  values.Get("backup_name").(string),
			BackupUUID:  values.Get("backup_uuid").(string),
			StartTime:   startTime,
			EndTime:     endTime,
			Status:      values.Get("status").(string),
			Msg:         values.Get("msg").(string),
			BackupInfos: backup,
		}

		backupinfos = append(backupinfos, backupinfo)
	}

	return backupinfos, nil
}
