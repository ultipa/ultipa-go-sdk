package test

import (
	"fmt"
	"log"
	"testing"
	"ultipa-go-sdk/utils"
)

func TestMerge(t *testing.T) {
	//node1String := `{"alias":"a","nodes":[{"id":100,"values":{"uuid":"546a2177bab956bdab434e5d681871bf","#khop_1":"3","type":"Human"}}]}`
	//node2String := `{"alias":"a","nodes":[{"id":200,"values":{"uuid":"546a2177bab956bdab434e5d681871bf","#khop_1":"3","type":"Human"}}]}`
	//var n1 interface{}
	//var n2 interface{}
	//_ = json.Unmarshal([]byte(node1String), &n1)
	//_ = json.Unmarshal([]byte(node2String), &n2)
	//dataMerge := utils.DataMerge{}
	//var nodes1 []*interface{}
	//var nodes2 []*interface{}
	//nodes1 = append(nodes1, &n1)
	//nodes2 = append(nodes2, &n2)
	//dataMerge.Concat(nodes1, nodes2, "alias")
}
func TestRegexp(t *testing.T)  {
	uqlString := `showTask().id( "1").name('abc').status(  'pengding'  ).limit(199 ).filter( {abc: "123", a: {$gt: 123}})`
	uql := utils.UQL{}
	uql.Parse(uqlString)
	fmt.Println(utils.StructToPrettyJSONString(uql))
}

func TestStringToTime(_ *testing.T) {
	ultipaTime := utils.UltipaTime{}

	d := "2021-03-19 12:04:05.123"
	t, err := ultipaTime.NewFromString(d)

	if err != nil {
		log.Fatalln(err)
	}

	utils.PrintJSON(t)

	// set uint64 to time
	t2 := ultipaTime.New(1849051732466851960)
	utils.PrintJSON(t2)

}

func CreateUltipaTime(t *testing.T) {
	//tt := utils.UltipaTime{}
	//tt.Update()
}