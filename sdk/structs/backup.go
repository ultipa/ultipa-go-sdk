package structs

type ShardInfo struct {
	Address    string `json:"address"`
	BackupPath string `json:"backup_path"`
	GraphID    int    `json:"graph_id"`
	ShardID    int    `json:"shard_id"`
	Type       string `json:"type"`
}

type BackupInfo struct {
	BackupName  string      `json:"backup_name"`
	BackupUUID  string      `json:"backup_uuid"`
	StartTime   string      `json:"start_time"`
	EndTime     string      `json:"end_time"`
	Status      string      `json:"status"`
	Msg         string      `json:"msg"`
	BackupInfos []ShardInfo `json:"backup_infos"`
}
