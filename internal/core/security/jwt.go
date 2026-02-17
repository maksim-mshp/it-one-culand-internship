package security

import "github.com/golang-jwt/jwt/v5"

func IsAdminRole(token string) (bool, error) {
	parser := jwt.NewParser()
	claims := jwt.MapClaims{}

	// TODO: validate secret key
	_, _, err := parser.ParseUnverified(token, claims)
	if err != nil {
		return false, err
	}

	role, _ := claims["role"].(string)
	return role == "admin", nil
}
