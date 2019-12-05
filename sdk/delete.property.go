package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

func deleteProperty(client ultipa.UltipaRpcsClient, _type ultipa.DeleteColumnRequest_DBType, name string) *ultipa.DeleteColumnReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.DeleteColumn(ctx, &ultipa.DeleteColumnRequest{
		Type:       _type,
		ColumnName: name,
	})

	if err != nil {
		log.Fatalf("[Error] delete property error: %v", err)
	}

	return msg
}

func DeleteNodeProperty(client ultipa.UltipaRpcsClient, name string) *ultipa.DeleteColumnReply {
	return deleteProperty(client, ultipa.DeleteColumnRequest_DBNODE, name)
}

func DeleteEdgeProperty(client ultipa.UltipaRpcsClient, name string) *ultipa.DeleteColumnReply {
	return deleteProperty(client, ultipa.DeleteColumnRequest_DBEDGE, name)
}
