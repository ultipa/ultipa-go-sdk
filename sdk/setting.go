package sdk

import (
	"context"
	"crypto/sha1"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

// GetSetting
func GetSetting(client Client, key string, username string) *ultipa.UserSettingReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, _ := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		Opt:      ultipa.UserSettingRequest_OPT_GET,
		Type:     key,
		UserName: username,
	})

	return res
}

// update Setting
func UpdateSetting(client Client, key string, value string, username string) *ultipa.UserSettingReply {
	ctx, err := context.WithTimeout(context.Background(), time.Second)

	if err != nil {
		log.Println("set setting ctx create failed")
	}

	res, _ := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		Opt:      ultipa.UserSettingRequest_OPT_SET,
		Type:     key,
		UserName: username,
		Data:     value,
	})

	// h := sha1.New()

	return res
}

func UpdateUser() {
	sha1.New()
}
