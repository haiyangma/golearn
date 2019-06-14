package mmh

import "github.com/twmb/murmur3"

func Hash(uid string) uint64 {
	return murmur3.Sum64([]byte(uid))
}