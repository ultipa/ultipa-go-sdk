package configuration

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

type PasswordEncryptor interface {
	Encrypt(string) string
}

var passwordEncryptMethodMap = make(map[string]PasswordEncryptor)

func init() {
	passwordEncryptMethodMap["MD5"] = MD5{}
	passwordEncryptMethodMap["LDAP"] = LDAP{}
	passwordEncryptMethodMap["NOTHING"] = NOTHING{}
}

func Encrypt(encryptorName string, pwd string) (string, error) {
	encryptorInterface, has := passwordEncryptMethodMap[strings.ToUpper(encryptorName)]
	if has {
		return encryptorInterface.Encrypt(pwd), nil
	}
	if encryptorName == "" {
		return passwordEncryptMethodMap["MD5"].Encrypt(pwd), nil
	}
	return "", errors.New(fmt.Sprintf("unsupported passpord encryptor"))
}

type MD5 struct{}
type LDAP struct{}
type NOTHING struct{}

func MD5Encrypt(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (t MD5) Encrypt(password string) string {
	return MD5Encrypt(password)
}

func (t LDAP) Encrypt(password string) string {
	return password
}

func (t NOTHING) Encrypt(password string) string {
	return password
}
