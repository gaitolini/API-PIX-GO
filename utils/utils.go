package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateTxID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
