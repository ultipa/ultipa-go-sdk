package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/printers"
)

func TestUQL(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	client, _ := GetClient([]string{"192.168.1.86:60072"}, "default")

	InitCases()

	for _, c := range cases {

		log.Println("Exec : ", c.UQL)

		//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
		resp, err := client.UQL(c.UQL, nil)

		if err != nil {
			log.Fatalln(err)
		}

		if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
			log.Println(resp.Status.Message)
			continue
		}

		for _, a := range resp.AliasList {
			dataitem := resp.Alias(a)
			printers.PrintAny(dataitem)
			log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

		}
	}
}

func TestUQL1(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	//client, _ := GetClient([]string{"192.168.1.86:60072"}, "default")
	//client, _ := GetClient([]string{"192.168.1.87:62061"}, "maker_test")
	//client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")
	//client, _ := GetClient([]string{"192.168.1.85:61115"}, "gongshang")
	client, _ := GetClient([]string{"192.168.1.87:60198"}, "ultipa_www")

	//uql := `n({@user && _uuid == 1}).e({@relation.relation_type == 'has'}).n({@projects} as project).re({@relation.relation_type == 'has'}).n({@etl} as etl) group by project skip 0 return table(project._id,project._uuid,count(etl)) as t limit 15 order by project.created_at desc`
	//uql := `find().edges(2658) as edges return edges{*}`
	//uql := `find().nodes() as nodes return nodes{*} limit 10`
	uql := `n({@version}).e().n({@docs_tree.status == 1} as book2).e().n({@lang.code == "en"}) with book2 as vbook
n(vbook).e().n({@docs_tree && is_root == "true" && @docs_tree.status == 1} as book1) with book1 as books
n(books).re({@docs_tree}).n(vbook).re({@docs_tree})[:3].n({@docs.status == 1 && @docs.type == "technical"|| @docs_tree} as n100) as path  with path as docs_path, n100 as books2
n({books || vbook || books2}).e({@docs_role || @docs_tree_role}).n({@role.name in ["public"]}) as p with p as role_path
return docs_path{*}, role_path{*},books`
	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
	}

	log.Println(resp.Statistic.TotalCost)

	//for _, a := range resp.AliasList {
	//	dataitem := resp.Alias(a)
	//	printers.PrintAny(dataitem)
	//	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//
	//}

}

func TestUQL2(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	//client, _ := GetClient([]string{"192.168.1.86:60072"}, "default")
	//client, _ := GetClient([]string{"192.168.1.87:62061"}, "maker_test")
	client, _ := GetClient([]string{"192.168.2.142:60062"}, "amz_zjs")

	//uql := `n({@user && _uuid == 1}).e({@relation.relation_type == 'has'}).n({@projects} as project).re({@relation.relation_type == 'has'}).n({@etl} as etl) group by project skip 0 return table(project._id,project._uuid,count(etl)) as t limit 15 order by project.created_at desc`
	uql := `find().nodes() return nodes limit 10`
	//uql := `find().nodes({@movie}) as nodes return nodes{*} limit 10`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintAny(resp.Get(0))
	nodes, _, _ := resp.Get(0).AsNodes()

	log.Println(nodes)

}

func TestUQL3(t *testing.T) {

	client, _ := GetClient([]string{"192.168.1.85:64801"}, "test_node_create8137")

	uql := `find().nodes({@test_schema2}) as n1 order by n1._uuid desc limit 1 return n1.test_timestamp`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}

	attrs, _ := resp.Get(0).AsAttr()

	for _, row := range attrs.Rows {
		log.Println(row)
	}

}

func TestUQL4(t *testing.T) {

	client, _ := GetClient([]string{"192.168.1.87:61095"}, "miniCircle")

	uql := `ab().src({_uuid == 1}).dest({_uuid == 3}).depth(:2) as paths with pnodes(paths) as nodeArray uncollect nodeArray as node return distinct(node)`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintAny(resp.Get(0))
}