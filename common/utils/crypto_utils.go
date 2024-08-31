package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword はパスワードをハッシュ化します
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// ComparePassword はハッシュ化されたパスワードと比較します
func ComparePassword(hashedPassword, plainPassword string) bool {
	return hashedPassword == HashPassword(plainPassword)
}
