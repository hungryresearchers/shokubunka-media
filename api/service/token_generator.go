package service

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateToken() string {
	b, err := GenerateRandomBytes()
	if err != nil {
		log.Fatal(err)
	}
	token := base64.URLEncoding.EncodeToString(b)
	return token
}

func GenerateRandomBytes() ([]byte, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
