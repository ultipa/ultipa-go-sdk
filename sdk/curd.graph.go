package sdk

import (
	"log"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

func (t *Connection) ListGraph(commonReq *types.Request_Common) {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_listGraph)

	res := t.UQLListSample(uql.ToString(), commonReq)
	log.Println(res.Status.Message, uql.ToString())
}
