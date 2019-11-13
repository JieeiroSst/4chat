package DataEncrypt

import (
	"crypto/sha256"
)
func HashData(data string) string {
	hash:=sha256.Sum256([]byte(data))
	
	return string(hash[:])
}
