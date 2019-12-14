package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type DeletePropertyResponse = ultipa.DeletePropertyReply

func deleteProperty(client ultipa.UltipaRpcsClient, _type ultipa.DBType, name string) *DeletePropertyResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.DeleteProperty(ctx, &ultipa.DeletePropertyRequest{
		Type:         _type,
		PropertyName: name,
	})

	if err != nil {
		log.Printf("[Error] delete property error: %v", err)
	}

	return msg
}

func DeleteNodeProperty(client ultipa.UltipaRpcsClient, name string) *DeletePropertyResponse {
	return deleteProperty(client, ultipa.DBType_DBNODE, name)
}

func DeleteEdgeProperty(client ultipa.UltipaRpcsClient, name string) *DeletePropertyResponse {
	return deleteProperty(client, ultipa.DBType_DBEDGE, name)
}
