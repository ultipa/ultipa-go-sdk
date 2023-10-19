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

	uql = ` grant().node_privilege(["READ"]).on("",@, *).user("lzq")`
	uqlItem = utils.NewUql(uql)
	isGlobal = uqlItem.IsGlobal()
	t.Logf("%s is global:%v", uql, isGlobal)

	uql = ` grant().edge_privilege(["READ"]).on("",@, *).user("lzq")`
	uqlItem = utils.NewUql(uql)
	isGlobal = uqlItem.IsGlobal()
	t.Logf("%s is global:%v", uql, isGlobal)


	uql = ` grant().privilege(["READ"]).on("",@, *).user("lzq")`
	uqlItem = utils.NewUql(uql)
	isGlobal = uqlItem.IsGlobal()
	t.Logf("%s is global:%v", uql, isGlobal)

	uql = ` grant().system().privilege(["STAT"]).user("lzq")`
	uqlItem = utils.NewUql(uql)
	isGlobal = uqlItem.IsGlobal()
	t.Logf("%s is global:%v", uql, isGlobal)
}
