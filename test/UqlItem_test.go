package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"testing"
)

func TestIsGlobalUql(t *testing.T) {
	uql := ` top().task("567")`
	uqlItem := utils.NewUql(uql)
	isGlobal := uqlItem.IsGlobal()
	t.Logf("%s is global:%v", uql, isGlobal)
}
