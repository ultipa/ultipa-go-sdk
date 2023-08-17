package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"log"
	"reflect"
	"testing"
)

//TODO:
func TestMergeStruct(t *testing.T) {
	type Company struct {
		Name string
	}

	type User struct {
		Name string
	}

	type Relations struct {
		Users     []*User
		Companies []*Company
	}

	R1 := &Relations{
		Users: []*User{
			{Name: "zhang"},
			{Name: "lin"},
		},
	}

	R2 := &Relations{
		Companies: []*Company{
			{Name: "Ultipa"},
			{Name: "Alibaba"},
		},
	}

	//E3 := struct {
	//
	//}{}

	log.Println(reflect.ValueOf(R1).Elem().Type().Field(0).Name)

	err := utils.MergeSameStruct(R1, R2)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v", R1)
}
