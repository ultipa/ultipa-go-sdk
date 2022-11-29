package test

import (
	"encoding/hex"
	"math/rand"
)

func RandStr(n int) string {
	result := make([]byte, n/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}


