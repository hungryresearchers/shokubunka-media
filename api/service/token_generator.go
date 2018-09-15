package service

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() (string, error) {
	b, err := GenerateRandomBytes()
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)
	return token, nil
}

func GenerateRandomBytes() ([]byte, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
