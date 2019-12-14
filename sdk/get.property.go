package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type Property struct {
	Name string
	Type PropertyType
}

func getPropertyInfo(client ultipa.UltipaRpcsClient, _type ultipa.DBType) []Property {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.GetPropertyInfo(ctx, &ultipa.GetPropertyInfoRequest{
		Type: _type,
	})

	var properties []Property
	for _, v := range msg.Properties {
		properties = append(properties, Property{
			Name: v.PropertyName,
			Type: v.PropertyType,
		})
	}

	if err != nil {
		log.Printf("[Error] get property error: %v", err)
	}

	return properties
}

func GetNodePropertyInfo(client ultipa.UltipaRpcsClient) []Property {
	return getPropertyInfo(client, ultipa.DBType_DBNODE)
}

func GetEdgePropertyInfo(client ultipa.UltipaRpcsClient) []Property {
	return getPropertyInfo(client, ultipa.DBType_DBEDGE)
}
