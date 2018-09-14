package service

import (
	"crypto/sha256"
)

func ToHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return string(hash[:])
}
