package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

func deleteProperty(client ultipa.UltipaRpcsClient, _type ultipa.DBType, name string) *ultipa.DeletePropertyReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.DeleteProperty(ctx, &ultipa.DeletePropertyRequest{
		Type:         _type,
		PropertyName: name,
	})

	if err != nil {
		log.Fatalf("[Error] delete property error: %v", err)
	}

	return msg
}

func DeleteNodeProperty(client ultipa.UltipaRpcsClient, name string) *ultipa.DeletePropertyReply {
	return deleteProperty(client, ultipa.DBType_DBNODE, name)
}

func DeleteEdgeProperty(client ultipa.UltipaRpcsClient, name string) *ultipa.DeletePropertyReply {
	return deleteProperty(client, ultipa.DBType_DBEDGE, name)
}
