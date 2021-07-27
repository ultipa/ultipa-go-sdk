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
		`algo(degree).params({})`,
		`alert().node_property()`,
		`n({_id == "C001"}).e().n({@card} as neighbors)
    find().nodes({_id == "C002"}) as C002
    with neighbors, C002
    update().nodes({_id == neighbors._id && balance > C002.balance}).set({level: level + 1})`,
	}


	for _, uql := range uqls {
		uqlParse := utils.EasyUqlParse{}
		uqlParse.Parse(uql)
		log.Printf("🧪 uql: %s", uqlParse.Uql)
		log.Printf("  HasWith: %t", uqlParse.HasWith())
		log.Printf("  HasWrite: %t", uqlParse.HasWrite())
		log.Printf("  HasAlgo: %t", uqlParse.HasAlgo())
		for _, item := range uqlParse.Commands {
			log.Printf(" -- %+v", item)
			log.Printf("    ---%+v【%d】", item.GetListParams(), len(item.GetListParams()))
		}
	}
}