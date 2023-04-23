package helpers

import (
	"fmt"

	"github.com/dg943/MyProject/backend/configs"
	jwt "github.com/dgrijalva/jwt-go"
)

func EncodeJWT(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret_key := []byte(configs.GetString("app_settings.jwt_key"))
	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret_key := []byte(configs.GetString("app_settings.jwt_key"))
		return secret_key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		// token.Claims is of type Claims and our helper package contains Claims type which is basically jwt.MapClaims, since it is implementing the Valid method
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(Claims)
	if !ok {
		return nil, fmt.Errorf("issue getting claims")
	}
	return &claims, nil
}
