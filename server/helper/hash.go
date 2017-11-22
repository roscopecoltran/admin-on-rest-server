package helper

import (
	"crypto/sha512"
	"encoding/hex"
)

func EncryptTextSHA512(text string) string {
	bytes := []byte(text)
	enc := sha512.Sum512(bytes)
	stringenc := hex.EncodeToString(enc[:])
	return stringenc
}
