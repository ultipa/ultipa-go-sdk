package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

type CreatePropertyResponse = ultipa.CreatePropertyReply

func createProperty(client ultipa.UltipaRpcsClient, dbType ultipa.DBType, propertyName string, propertyType PropertyType) *CreatePropertyResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.CreateProperty(ctx, &ultipa.CreatePropertyRequest{
		Type: dbType,
		Properties: []*ultipa.CreatePropertyValues{
			&ultipa.CreatePropertyValues{
				PropertyName: propertyName,
				PropertyType: propertyType,
			},
		},
	})

	if err != nil {
		log.Printf("[Error] create node property error: %v", err)
	}

	return msg
}

func CreateNodeProperty(client ultipa.UltipaRpcsClient, propertyName string, propertyType PropertyType) *CreatePropertyResponse {
	return createProperty(client, ultipa.DBType_DBNODE, propertyName, propertyType)
}

func CreateEdgeProperty(client ultipa.UltipaRpcsClient, propertyName string, propertyType PropertyType) *CreatePropertyResponse {
	return createProperty(client, ultipa.DBType_DBEDGE, propertyName, propertyType)
}
