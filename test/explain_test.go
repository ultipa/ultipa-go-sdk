package test

import (
	"log"
	"testing"
	"ultipa-go-sdk/sdk/printers"
)

func TestExplain(t *testing.T) {

	client, _ := GetClient([]string{"192.168.1.85:54001"}, "default")

	//resp, err := client.UQL("explain find().nodes() as nodes limit 1 return nodes limit 10", nil)
	//	resp, err := client.UQL(`explain n({@account} as buyer).e().n({@card}).re({@transaction} as buy).n()
	//with buyer, buy
	//group by buyer
	//with collect(distinct(day_of_week(buy.time))) as array
	//where (1 in array || 7 in array) && 2 nin array && 3 nin array && 4 nin array && 5 nin array && 6 nin array
	//return buyer{*} limit 100`, nil)

	resp, err := client.UQL(`explain find().nodes() as n1 limit 1 
 find().nodes() as n2 limit 1
 find().nodes() as n3 limit 1
return n1,n2,n3 limit 10`, nil)

	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(resp)
	explain := resp.ExplainPlan.Explain
	printers.PrintExplain(explain)

}
