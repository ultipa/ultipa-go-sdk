package test

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"testing"
)

func TestIsNeedToEscapeName(t *testing.T) {
	cases := map[string]bool{
		"`ab@cd`":  false,
		"abc@d":    false,
		"`abcd":    false,
		"abcd`":    false,
		"`abcd`":   false,
		"哈哈abcd": true,
		"abcd\"":   true,
	}

	for name, expected := range cases {
		actual := utils.IsNeedToEscapeName(name)
		if actual != expected {
			t.Errorf("%v expected: %v, actual:%v", name, expected, actual)
		}

	}

}
