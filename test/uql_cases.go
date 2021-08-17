package test

type Case struct {
	UQL   string
	Alias []string
	Type  string
}

var cases []*Case

func InitCases() {

	// init single returns
	cases = append(cases,
		//&Case{
		//	UQL:   "n().e().n(as end) return end{*} as end1 limit 10;",
		//	Alias: []string{"end1"},
		//},
		//&Case{
		//	UQL:   "find().nodes() as nodes return nodes limit 10",
		//	Alias: []string{"nodes"},
		//},
		&Case{
			UQL:   "show().graph()",
			Alias: []string{"_graph"},
		},
		//&Case{
		//	UQL:   "show().schema()",
		//	Alias: []string{"_nodeSchema", "_edgeSchema"},
		//	Type:  "schema",
		//},
		//&Case{
		//	UQL:   "n().e().n() as paths return paths limit 10;",
		//	Alias: []string{"paths"},
		//},
		//&Case{
		//	UQL:   "show().index()",
		//	Alias: []string{"nodeIndex","edgeIndex"},
		//},
		//&Case{
		//	UQL:   "show().graph(\"multi_schema_test\")",
		//	Alias: []string{"_graph"},
		//},
		//&Case{
		//	UQL:   "find().nodes({@amz}) as nodes return nodes.name as name limit 10",
		//	Alias: []string{"name"},
		//},
		//&Case{
		//	UQL:   "find().nodes({@amz}).limit(10) as nodes return collect(nodes.name) as name ",
		//	Alias: []string{"name"},
		//},
		//&Case{
		//	UQL:   "find().nodes({@amz}) as nodes return count(nodes) as totalName;",
		//	Alias: []string{"totalName"},
		//
		//&Case{
		//	UQL:   "find().nodes({@amz}) as nodes return count(nodes) as totalName;",
		//	Alias: []string{"totalName"},
		//},


	)
}
