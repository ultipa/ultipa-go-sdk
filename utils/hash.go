package utils

import "hash/fnv"

func Hash64(str []byte) int64 {
	h := fnv.New64a()
	h.Write(str)
	return int64(h.Sum64())
}
