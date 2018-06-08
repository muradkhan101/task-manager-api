package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	SECRET_KEY = "poopydoopydoo"
	DATA       = "*** this is a secret ***"
)

func main() {
	h := hmac.New(sha256.New, []byte(SECRET_KEY))
	h.Write([]byte(DATA))
	str := h.Sum(nil)
	fmt.Println("BYTE ARRAY:  ", str)
	fmt.Println("STRINGIFIED:  ", hex.EncodeToString(str))
}
