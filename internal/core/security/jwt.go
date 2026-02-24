package security

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func IsAdminRole(tokenString, secretKey string) (bool, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)

	claims := jwt.MapClaims{}

	token, err := parser.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Printf("jwt validation error: %v", err)
		return false, err
	}
	if !token.Valid {
		return false, err
	}

	role, _ := claims["role"].(string)
	return role == "admin", nil
}
