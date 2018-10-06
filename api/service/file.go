package service

import (
	"crypto/rand"
	"encoding/hex"
	"path/filepath"
)

func ChangeUniqueName(previousName string) string {
	suffix := filepath.Ext(previousName)
	nextName := randomName()
	return nextName + suffix
}

func randomName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}
