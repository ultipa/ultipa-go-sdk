package test

type Case struct {
	UQL string
	Alias []string
	Type string
}


var cases []*Case


func InitCases() {

	// init single returns
	cases = append(cases, &Case{
		UQL: "n().e().n(as end) return end{*} limit 10;",
		Alias: []string{"end"},
	}, &Case{
		UQL: "show().graph()",
		Alias :[]string{"graph"},
	}, &Case{
		UQL: "show().schema()",
		Alias: []string{"nodeSchema", "edgeSchema"},
		Type: "schema",
	})
}


