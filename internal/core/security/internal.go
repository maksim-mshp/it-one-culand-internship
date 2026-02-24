package security

import "crypto/subtle"

func IsValidInternal(tokenString, secretKey string) bool {
	return subtle.ConstantTimeCompare([]byte(tokenString), []byte(secretKey)) == 1
}
