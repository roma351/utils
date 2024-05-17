package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"math/big"
	"strings"
)

func RandomString(length int) string {
	source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			log.Println(err)
		}
		result[i] = source[nBig.Int64()]
	}

	return string(result)
}

func RandomStringInt(length int) string {
	source := "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			log.Println(err)
		}
		result[i] = source[nBig.Int64()]
	}

	return string(result)
}

func RandomStringAndInt(length int) string {
	source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			log.Println(err)
		}
		result[i] = source[nBig.Int64()]
	}

	return string(result)
}

func Cut(s string, v1 string, v2 string) (string, error) {
	sa := strings.Split(s, v1)
	if len(sa) > 1 {
		sa := strings.Split(sa[1], v2)
		if len(sa) > 0 {
			return sa[0], nil
		}
	}
	return "", errors.New("ERROR 2222")
}

func StringInSlice(target string, slice []string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func IntInSlice(target int, slice []int) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func Int32InSlice(target int32, slice []int32) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func Int64InSlice(target int64, slice []int64) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func IOToString(value io.Reader) (string, error) {
	b, err := io.ReadAll(value)
	return string(b), err
}

func IOToStringNoErr(value io.Reader) string {
	b, err := io.ReadAll(value)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func GenerateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
