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
		//&Case{
		//	UQL:   "show().graph()",
		//	Alias: []string{"graph"},
		//},
		//&Case{
		//	UQL:   "show().schema()",
		//	Alias: []string{"nodeSchema", "edgeSchema"},
		//	Type:  "schema",
		//},
		&Case{
			UQL:   "n().e().n() as paths return paths limit 10;",
			Alias: []string{"paths"},
		},
	)
}
