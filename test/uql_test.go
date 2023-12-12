package test

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"log"
	"testing"
)

func TestUQL(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	client, _ := GetClient(hosts, graph)

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
	client, _ := GetClient(hosts, graph)

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
	client, _ := GetClient(hosts, graph)

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

	client, _ := GetClient(hosts, graph)

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

	client, _ := GetClient(hosts, graph)

	uql := `ab().src({_uuid == 1}).dest({_uuid == 3}).depth(:2) as paths with pnodes(paths) as nodeArray uncollect nodeArray as node return distinct(node)`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintAny(resp.Get(0))
}

func TestUQL5(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := `find().nodes({@account.year==1978 && @account.name=="念敏"}) as nodes return nodes{*} limit 1`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintAny(resp.Get(0))
}

func TestUQL6(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "find().nodes({@movie}) as nodes return table(nodes.timestamp,nodes.frating) limit 0"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
	table, err := resp.Alias("table(nodes.timestamp,nodes.frating)").AsTable()
	if err != nil {
		return
	}
	printers.PrintTable(table)
}

func TestUQLAlterGraph(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "alter().graph('alter_graph_1').set({name:'alter_graph'})" //test123
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
}

func TestUQLCompactWithNotExistGraph(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := `compact().graph("c1")`
	//ty, leader, follower, global, err := client.GetConnByUQL(uql, "random_test_js_1672976970614")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//t.Logf("uql type:%v", ty)
	//t.Logf("leader:%s", leader.Host)
	//sb := strings.Builder{}
	//for _, f := range follower {
	//	sb.WriteString(f.Host)
	//	sb.WriteString(",")
	//}
	//t.Logf("followers:%s", sb.String())
	//
	//t.Logf("global leader:%s", global.Host)

	resp, err := client.UQL(uql, &configuration.RequestConfig{
		GraphName: "c1",
	})
	if err != nil {
		t.Fatalf("fail to compact:%v", err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Fatalf(resp.Status.Message)
	}
}

func TestTopUql(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := `top()`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintAny(resp.Get(0))
}
func TestUQLFindNodesWithList(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "find().nodes({@People}) as nodes return nodes{*}"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
	nodes, schemas, err := resp.Alias("nodes").AsNodes()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintNodes(nodes, schemas)
}

func TestUQLFindNodesWithAttrList(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "find().nodes() as nodes return collect(distinct(nodes)) as arrNode"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
	attrs, err := resp.Alias("arrNode").AsAttr()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintAttr(attrs)
}

func TestUQLFindNodesWithAttrListNullValue(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "find().nodes({@nodeSchemaList}) as n return collect(n.stringList)"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
	attrs, err := resp.Alias("collect(n.stringList)").AsAttr()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintAttr(attrs)
}

func TestUQLFindNodesAsAttrList(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "find().nodes({_uuid < 10}) return collect(nodes)"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
	attr, err := resp.Alias("collect(nodes)").AsAttr()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintAttr(attr)
}

func TestUQLFindPathsWithGroupByAttr(t *testing.T) {

	client, _ := GetClient(hosts, graph)

	uql := "n({_uuid in [4,5,6]} as n1).e().n(as n2) as paths group by n1 return n1{*}, collect(paths)"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
	//断言返回数据
	node, schemas, err := resp.Alias("n1").AsNodes()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintNodes(node, schemas)

	attr, err := resp.Alias("collect(paths)").AsAttr()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintAttr(attr)
}

func TestUqlInsertListProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	uql := `insert().nodes({name:["zhangsan","lisi"]}).into(@People)`
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
}

func TestUQLWithLimit(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	uql := "find().edges({@insertEdge}) as edges return edges{*} limit 40"
	//uql := "find().nodes({@insertNode}) as n return n{typeListString,typeListInt32,typeListInt64,typeListUint32,typeListUint64,typeListFloat,typeListDouble,typeListDatetime,typeListTimestamp,typeListText}"
	//uql := "find().nodes({@insertNode}) as nodes return nodes{*}"
	resp, _ := client.UQL(uql, nil)
	edges, e, _ := resp.Alias("edges").AsEdges()

	printers.PrintEdges(edges, e)
	//var nodes1 []*structs.Node
	//nodes1 = append(nodes1, nodes[0])
	//fmt.Println(nodes1)
}

func TestUqlFindPointProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	uql := `find().nodes([11,12]) as nodes return nodes.typePoint`
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
	attr, err := resp.Alias("nodes.typePoint").AsAttr()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintAttr(attr)
}

func TestOnePathAsPaths(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	var uql = "ab().src(51).dest(103).depth(1) as paths return paths{}"
	resp, _ := client.UQL(uql, nil)
	paths, _ := resp.Alias("paths").AsPaths()
	printers.PrintPaths(paths)
}
func TestDateTime(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	var uql = "find().nodes({@`nodeSchema3`}) as nodes return table(nodes.typeListDatetime[0])"
	resp, _ := client.UQL(uql, nil)
	table, _ := resp.Alias("table(nodes.typeListDatetime[0])").AsTable()
	printers.PrintTable(table)
}

func TestUqlFindWithDecimalProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	uql := `find().nodes({@default}) as nodes return nodes{*}`
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
	nodes, schemas, err := resp.Alias("nodes").AsNodes()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintNodes(nodes, schemas)
}

func TestUqlFindWithSetProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	uql := `find().nodes({@default}) as nodes return nodes{*}`
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
	nodes, schemas, err := resp.Alias("nodes").AsNodes()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintNodes(nodes, schemas)
}

func TestFindNodeWithOptionalUql(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	uql := "OPTIONAL find().nodes({@account.year < 1969}) as nodes return nodes{*}"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
	nodes, schemas, err := resp.Alias("nodes").AsNodes()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintNodes(nodes, schemas)
}
func TestFindEdgeWithOptionalUql(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	uql := "OPTIONAL find().edges({@disagree.targetPost <10}) as edges return edges{*}"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
	edges, schemas, err := resp.Alias("edges").AsEdges()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintEdges(edges, schemas)
}

func TestFindPathWithOptionalUql(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	uql := "OPTIONAL n(1).e().n(103) as paths RETURN paths{*}"
	resp, err := client.UQL(uql, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		t.Fatal(resp.Status.Message)
	}
	paths, err := resp.Alias("paths").AsPaths()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintPaths(paths)
}
