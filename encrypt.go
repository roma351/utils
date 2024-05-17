package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func AESEncrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGcm.Seal(nil, nonce, plaintext, nil)
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	macSum := mac.Sum(nil)

	finalData := append(nonce, append(ciphertext, macSum...)...)
	return base64.StdEncoding.EncodeToString(finalData), nil
}

func AESDecrypt(encodedData string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}

	nonceSize := 12 // Размер nonce для AES-GCM
	if len(data) < nonceSize+sha256.Size {
		return nil, errors.New("invalid data size")
	}

	nonce, ciphertext, macSum := data[:nonceSize], data[nonceSize:len(data)-sha256.Size], data[len(data)-sha256.Size:]

	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	expectedMacSum := mac.Sum(nil)

	if !hmac.Equal(macSum, expectedMacSum) {
		return nil, errors.New("invalid MAC")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func AESKeyGenerate() ([]byte, error) {
	return KeyGenerate(32)
}
