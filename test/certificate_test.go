package test

import (
	"fmt"
	"testing"
	"ultipa-go-sdk/sdk/utils"
)

func TestGetCertificate(t *testing.T) {

	certificate := utils.GetCertificate("kinqhpwws.us-east-2.uct.ultipa-inc.org:60010")

	fmt.Println(certificate)
}
