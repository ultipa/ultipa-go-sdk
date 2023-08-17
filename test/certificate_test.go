package test

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"testing"
)

func TestGetCertificate(t *testing.T) {

	certificate := utils.GetCertificate("kinqhpwws.us-east-2.uct.ultipa-inc.org:60010")

	fmt.Println(certificate)
}
