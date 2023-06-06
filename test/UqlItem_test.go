package test

import (
	"testing"
	"ultipa-go-sdk/sdk/utils"
)

func TestIsGlobalUql(t *testing.T) {
	uql := ` top().task("567")`
	uqlItem := utils.NewUql(uql)
	isGlobal := uqlItem.IsGlobal()
	t.Logf("%s is global:%v", uql, isGlobal)
}
