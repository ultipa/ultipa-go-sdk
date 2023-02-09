package test

import (
	"fmt"
	"log"
	"regexp"
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

func TestIsExtra(t *testing.T) {

	uqls := map[string]bool{}
	for uql, _ := range utils.ExtraUqlCommandKeys {
		uqls[uql] = true
	}

	nonExtraUql := []string{
		`find().nodes({(_id == 1 && c == 2)})`,
		`create().graph("<name>", "<desc?>")`,
		`exec task algo(degree).params({})`,
		"show().edge_schema(@amz).limit(100)",
		"show().edge_schema(@amz) limit 10",
	}
	for _, s := range nonExtraUql {
		uqls[s] = false
	}

	for uql, isExtra := range uqls {
		Uql := utils.NewUql(uql)
		actual := Uql.IsExtra()
		log.Printf("🧪 uql: %s", Uql.Uql)
		log.Printf("  isExtra: %t", actual)
		log.Println("-------")
		if isExtra != actual {
			t.Errorf("%s is extra? expected:%t,actual:%t", uql, isExtra, actual)
		}
	}
}

func TestRegularExpress(t *testing.T) {
	uqls := []string{
		"find().nodes({(_id == 1 && c == 2)})",
		"show().graph()",
		`show().graph("name")`,
		`show().graph(name)`,
		`create().graph("<name>", "<desc?>")`,
		`algo(degree).params({})`,
		`exec task algo(degree).params({})`,
		`alert().node_property()`,
		`n({_id == "C001"}).e().n({@card} as neighbors)
	find().nodes({_id == "C002"}) as C002
	with neighbors, C002
	update().nodes({_id == neighbors._id && balance > C002.balance}).set({level: level + 1})`,
		"",
		"show().edge_schema(@amz).limit(100)",
		"show().edge_schema(@amz) limit 10",
		"top()",
		"kill()",
		"grant().user",
	}
	//matcher := regexp.MustCompile(`([a-z_A-Z]*)\(([^\(|^\)]*)\)`)
	matcher := regexp.MustCompile(`([a-z_A-Z]*)(?:\((?:[^\(|^\)]*)\))?(?:[.]*([a-z_A-Z]*))*`)
	for idx, uql := range uqls {
		fmt.Printf("%d:%q\n", idx, matcher.FindStringSubmatch(uql))
	}
}
