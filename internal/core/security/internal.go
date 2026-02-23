package security

func IsValidInternal(tokenString, secretKey string) bool {
	return tokenString == secretKey
}
