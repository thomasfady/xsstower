package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHex(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
