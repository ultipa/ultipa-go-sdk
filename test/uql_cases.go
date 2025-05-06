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
		//	Uql:   "n().e().n(as end) return end{*} as end1 limit 10;",
		//	Alias: []string{"end1"},
		//},
		//&Case{
		//	Uql:   "find().nodes() as nodes return nodes limit 10",
		//	Alias: []string{"nodes"},
		//},
		//&Case{
		//	Uql:   "exec task show().graph()",
		//	Alias: []string{"_graph"},
		//},
		//&Case{
		//	Uql:   "show().schema()",
		//	Alias: []string{"_nodeSchema", "_edgeSchema"},
		//	DBType:  "schema",
		//},
		//&Case{
		//	Uql:   "n().e().n() as paths return paths limit 10;",
		//	Alias: []string{"paths"},
		//},
		//&Case{
		//	Uql:   "show().index()",
		//	Alias: []string{"nodeIndex","edgeIndex"},
		//},
		//&Case{
		//	Uql:   "show().graph(\"multi_schema_test\")",
		//	Alias: []string{"_graph"},
		//},
		//&Case{
		//	Uql:   "find().nodes({@amz}) as nodes return nodes.name as name limit 10",
		//	Alias: []string{"name"},
		//},
		//&Case{
		//	Uql:   "find().nodes({@amz}).limit(10) as nodes return collect(nodes.name) as name ",
		//	Alias: []string{"name"},
		//},
		//&Case{
		//	Uql:   "find().nodes({@amz}) as nodes return count(nodes) as totalName;",
		//	Alias: []string{"totalName"},
		//
		//&Case{
		//	Uql:   "find().nodes() as nodes return count(distinct(nodes.name)) as totalName;",
		//	Alias: []string{"totalName"},
		//},
		//&Case{
		//	Uql:   "UFE().node_property(@User.name)",
		//	Alias: []string{},
		//},
		//&Case{
		//	Uql:   "LTE().node_property(@User.name)",
		//	Alias: []string{},
		//},
		//&Case{
		//	Uql:   "show().schema()",
		//	Alias: []string{},
		//},
		&Case{
			UQL:   "find().nodes() return nodes limit 1",
			Alias: []string{},
		},
		&Case{
			UQL:   "find().nodes() return nodes limit 1",
			Alias: []string{},
		},
		&Case{
			UQL:   "find().nodes() return nodes limit 1",
			Alias: []string{},
		},
		&Case{
			UQL:   "find().nodes() return nodes limit 1",
			Alias: []string{},
		},
	)
}
