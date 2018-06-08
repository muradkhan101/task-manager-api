package backend

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

func Encrypt(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func CompareEncoded(expected string, actual string) bool {
	exp, _ := hex.DecodeString(expected)
	act, _ := hex.DecodeString(actual)
	return hmac.Equal(exp, act)
}

func MakeSalt(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(21, 125))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return rand.Intn(max) + min
}
