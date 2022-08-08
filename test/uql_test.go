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
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	//uql := `n({@user && _uuid == 1}).e({@relation.relation_type == 'has'}).n({@projects} as project).re({@relation.relation_type == 'has'}).n({@etl} as etl) group by project skip 0 return table(project._id,project._uuid,count(etl)) as t limit 15 order by project.created_at desc`
	uql := `find().edges(2658) as edges return edges{*}`
	//uql := `find().nodes({@movie}) as nodes return nodes{*} limit 10`

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

func TestUQL2(t *testing.T) {

	//client, _ := GetClient([]string{"210.13.32.146:40101"}, "default")
	//client, _ := GetClient([]string{"192.168.1.94:60061"}, "default")
	//client, _ := GetClient([]string{"192.168.1.86:60072"}, "default")
	//client, _ := GetClient([]string{"192.168.1.87:62061"}, "maker_test")
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")

	//uql := `n({@user && _uuid == 1}).e({@relation.relation_type == 'has'}).n({@projects} as project).re({@relation.relation_type == 'has'}).n({@etl} as etl) group by project skip 0 return table(project._id,project._uuid,count(etl)) as t limit 15 order by project.created_at desc`
	uql := `n(2).e()[:2].n(459) as path return pnodes(path) as kk`
	//uql := `find().nodes({@movie}) as nodes return nodes{*} limit 10`

	log.Println("Exec : ", uql)

	//resp, err := client.UQL(c.UQL, &configuration.RequestConfig{GraphName: "multi_schema_test"})
	resp, err := client.UQL(uql, nil)

	if err != nil {
		log.Fatalln(err)
	}

	array, err := resp.Alias("kk").AsArray()

	printers.PrintArray(array)

}
