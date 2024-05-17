package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
