package http

//_graph
//_nodeSchema
//_edgeSchema
//_nodeProperty
//_edgeProperty
//_nodeIndex
//_edgeIndex
//_nodeFulltext
//_edgeFulltext
//_statistic
//_top
//_task
//_policy
//_user
//_privilege

const (
	RESP_GRAPH_KEY         string = "_graph"
	RESP_NODE_SCHEMA_KEY   string = "_nodeSchema"
	RESP_EDGE_SCHEMA_KEY   string = "_edgeSchema"
	RESP_GRAPH_COUNT_KEY   string = "_graphCount" // 5.0 showSchema get total
	RESP_NODE_PROPERTY_KEY string = "_nodeProperty"
	RESP_EDGE_PROPERTY_KEY string = "_edgeProperty"
	RESP_NODE_INDEX_KEY    string = "_nodeIndex"
	RESP_EDGE_INDEX_KEY    string = "_edgeIndex"
	RESP_NODE_FULLTEXT_KEY string = "_nodeFulltext"
	RESP_EDGE_FULLTEXT_KEY string = "_edgeFulltext"
	RESP_STATISTIC_KEY     string = "_statistic"
	RESP_TOP_KEY           string = "_top"
	RESP_TASK_KEY          string = "_task" // asTask
	RESP_POLICY_KEY        string = "_policy"
	RESP_USER_KEY          string = "_user"
	RESP_PRIVILEGE_KEY     string = "_privilege"
	RESP_ALGOS_KEY         string = "_algoList"
	RESP_EXTAS_KEY         string = "_extaList"
	RESP_JOB_KEY           string = "result"
	RESP_LICENSE_KEY       string = "license"
	RESP_PROJECT_KEY       string = "_projectList"
)
