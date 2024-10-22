package api

func (api *UltipaAPI) GetServerVersion() (string, error) {
	//resp, err := api.Uql("stats()", nil)
	//if err != nil {
	//	return "", err
	//}
	//if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
	//	return "", errors.New(resp.Status.Message)
	//}
	//stat, err := resp.Alias(http.RESP_STATISTIC_KEY).AsTable()
	//if err != nil {
	//	return "", err
	//}
	//rows := stat.ToKV()
	//if len(rows) > 0 {
	//	serverVersion := rows[0].Get("version")
	//	if serverVersion != nil {
	//		serverVer := serverVersion.(string)
	//		return serverVer, nil
	//	}
	//}
	//return "", nil

	//stats, err := api.Stats(nil)
	//if err != nil {
	//	return "", err
	//}
	//return stats.Version, nil
	return "", nil
}
