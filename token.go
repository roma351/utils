package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func KeyGenerate(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func EncodeToken(token []byte) string {
	return base64.StdEncoding.EncodeToString(token)
}

func DecodeToken(token string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(token)
}
