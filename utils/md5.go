package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// lowercase
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// uppercase
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// Encode
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// Decode
func ValidPassword(plainpwd, salt string, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
