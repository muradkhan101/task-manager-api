package backend

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func Encrypt(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func CompareEncoded(expected string, actual string, secret string) bool {
	act, _ := hex.DecodeString(actual)
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(expected))
	return hmac.Equal(h.Sum(nil), act)
}

func MakeSalt(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(25, 125))
	}
	return hex.EncodeToString(bytes)
}

func randInt(min int, max int) int {
	return rand.Intn(max) + min
}
