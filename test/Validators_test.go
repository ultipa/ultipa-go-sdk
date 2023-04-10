package test

import (
	"testing"
	"ultipa-go-sdk/sdk/utils"
)

func TestCheckIsEscapedName(t *testing.T) {
	cases := map[string]bool{
		"`ab@cd`":  true,
		"abc@d":    false,
		"`abcd":   false,
		"abcd`":   false,
		" `abcd`": false,
	}

	for name, expected := range cases {
		actual := utils.CheckIsEscapedName(name)
		if actual != expected {
			t.Errorf("%v expected: %v, actual:%v", name, expected, actual)
		}

	}

}
