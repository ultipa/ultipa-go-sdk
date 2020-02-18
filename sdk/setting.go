package sdk

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
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

// SetUser
func SetUser(client Client, username string, password string, hosts []string) (*ultipa.UserSettingReply, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(password))
	macPass := mac.Sum(nil)

	// log.Println(base64.StdEncoding.EncodeToString(macPass))

	expire := (time.Now().UnixNano() + int64(time.Second*60*60*24*30)) / int64(time.Millisecond)

	data, errj := json.Marshal(map[string]interface{}{
		"username": username,
		"password": base64.StdEncoding.EncodeToString(macPass),
		"isAdmin":  false,
		"hosts":    hosts,
		"expire":   expire,
	})

	if errj != nil {

		log.Fatal(errj)
	}

	res, err := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		Opt:      ultipa.UserSettingRequest_OPT_SET,
		Type:     userinfoKey,
		UserName: username,
		Data:     string(data),
	})

	// log.Printf("\n%v\n", string(data))

	return res, err
}
