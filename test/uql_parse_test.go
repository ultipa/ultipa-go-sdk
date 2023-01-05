package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/sdk/utils"
)

func TestEasyUqlParse(t *testing.T) {
	uqls := []string{
		"find().nodes({_id == 1 && c == 2})",
		"show().graph()",
		`show().graph("name")`,
		`show().graph(name)`,
		`create().graph("<name>", "<desc?>")`,
		`exec task algo(degree).params({})`,
		`alter().node_property()`,
		`n({_id == "C001"}).e().n({@card} as neighbors)
    find().nodes({_id == "C002"}) as C002
    with neighbors, C002
    update().nodes({_id == neighbors._id && balance > C002.balance}).set({level: level + 1})`,
	}

	for _, uql := range uqls {
		Uql := utils.NewUql(uql)
		log.Printf("🧪 uql: %s", Uql.Uql)
		log.Printf("  HasWith: %t", Uql.HasWith())
		log.Printf("  HasWrite: %t", Uql.HasWrite())
		log.Printf("  HasTask: %t", Uql.HasExecTask())
		log.Printf("  IsGlobal: %t", Uql.IsGlobal())
	}
}

func TestParseGraph(t *testing.T) {
	uqls := []string{
		`mount( ).graph("abcde")`,
		`mount( ).graph('abcde')`,
		`unmount( ).graph("abcde")`,
		`unmount( ).graph('abcde')`,
		`truncate().graph("abcde")`,
		`truncate().graph('abcde')`,
		`show().graph(name)`,
		`create().graph("abcde", "desc")`,
		`exec task algo(degree).params({})`,
		`alter().node_property()`,
	}

	for idx, uql := range uqls {
		Uql := utils.NewUql(uql)
		ok, graph := Uql.ParseGraph()
		t.Logf("uql[%d]:%v, graph:%s", idx, ok, graph)
	}

}
