package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type PropertyType = ultipa.UltipaPropertyType

const (
	PROPERTY_TYPE_INT     PropertyType = ultipa.UltipaPropertyType_PROPERTY_INT
	PROPERTY_TYPE_STRING  PropertyType = ultipa.UltipaPropertyType_PROPERTY_STRING
	PROPERTY_TYPE_TEXT    PropertyType = ultipa.UltipaPropertyType_PROPERTY_TEXT
	PROPERTY_TYPE_BOOLEAN PropertyType = ultipa.UltipaPropertyType_PROPERTY_BOOLEAN
	PROPERTY_TYPE_UNKNOWN PropertyType = ultipa.UltipaPropertyType_PROPERTY_UNKNOWN
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
