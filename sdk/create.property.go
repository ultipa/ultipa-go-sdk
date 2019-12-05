package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type PropertyType ultipa.UltipaColumnType

const (
	PROPERTY_TYPE_INT     ultipa.UltipaColumnType = ultipa.UltipaColumnType_COLUMN_INT
	PROPERTY_TYPE_STRING  ultipa.UltipaColumnType = ultipa.UltipaColumnType_COLUMN_STRING
	PROPERTY_TYPE_UNKNOWN ultipa.UltipaColumnType = ultipa.UltipaColumnType_COLUMN_UNKNOWN
)

// type PropertyType struct {
// 	STRING ultipa.UltipaColumnType_COLUMN_STRING
// 	INT ultipa.UltipaColumnType_COLUMN_INT
// 	UNKNOWN ultipa.UltipaColumnType_COLUMN_UNKNOWN
// }

func createProperty(client ultipa.UltipaRpcsClient, dbType ultipa.CreatePropertyRequest_DBType, propertyName string, propertyType ultipa.UltipaColumnType) *ultipa.CreatePropertyReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.CreateProperty(ctx, &ultipa.CreatePropertyRequest{
		Type: dbType,
		Properties: []*ultipa.CreatePropertyValues{
			&ultipa.CreatePropertyValues{
				ColumnName: propertyName,
				ColumnType: propertyType,
			},
		},
	})

	if err != nil {
		log.Fatalf("[Error] create node property error: %v", err)
	}

	return msg
}

func CreateNodeProperty(client ultipa.UltipaRpcsClient, propertyName string, propertyType ultipa.UltipaColumnType) *ultipa.CreatePropertyReply {
	return createProperty(client, ultipa.CreatePropertyRequest_DBNODE, propertyName, propertyType)
}

func CreateEdgeProperty(client ultipa.UltipaRpcsClient, propertyName string, propertyType ultipa.UltipaColumnType) *ultipa.CreatePropertyReply {
	return createProperty(client, ultipa.CreatePropertyRequest_DBEDGE, propertyName, propertyType)
}
