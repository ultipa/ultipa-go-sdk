package structs

type Top struct {
	ProcessId  string `json:"process_id"`
	Status     string `json:"status"`
	ProcessUql string `json:"process_uql"`
	Duration   string `json:"duration"`
}
