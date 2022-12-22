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
	client, _ := GetClient([]string{"192.168.1.85:61115"}, "gongshang")

	//uql := `n({@user && _uuid == 1}).e({@relation.relation_type == 'has'}).n({@projects} as project).re({@relation.relation_type == 'has'}).n({@etl} as etl) group by project skip 0 return table(project._id,project._uuid,count(etl)) as t limit 15 order by project.created_at desc`
	//uql := `find().edges(2658) as edges return edges{*}`
	uql := `find().nodes() as nodes return nodes{*} limit 10`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
	}

	for _, a := range resp.AliasList {
		dataitem := resp.Alias(a)
		printers.PrintAny(dataitem)
		log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

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

func TestUQL5(t *testing.T) {

	client, _ := GetClient([]string{"192.168.1.87:61095"}, "miniCircle")

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

	client, _ := GetClient([]string{"192.168.1.87:61095"}, "miniCircle")

	uql := `find().edges({_uuid in [150,449]}) return edges{*}`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}
	alias := resp.Alias("edges")
	//printers.PrintAny(alias)

	nodes, schema, err := alias.AsEdges()
	printers.PrintEdges(nodes, schema)

}
