package sdk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
)

const (
	secretKey   string = "ultipa"
	userinfoKey string = "user_info"
)

func Auth(client Client, username string, password string) bool {

	userinfo := GetSetting(client, userinfoKey, username)

	// fmt.Println(userinfo)

	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(password))
	macPass := mac.Sum(nil)
	// fmt.Println(username, password, macPass)

	var user map[string]interface{}

	json.Unmarshal([]byte(userinfo.Data), &user)

	// fmt.Println(user)

	if user["password"] == base64.StdEncoding.EncodeToString(macPass) {
		return true
	}
	return false
}
