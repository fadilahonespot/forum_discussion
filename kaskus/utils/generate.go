package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)


func GenerateToken(id int, pass string) (string, error) {
	claims := jwt.MapClaims{}

	claims["user_id"] = id
	claims["pass"] = pass

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("auth.secret")))
	if err != nil {
		fmt.Printf("[utils.GenerateToken] Error to signature generate token, %v \n", err)
		return "", fmt.Errorf("Failed generate token")
	}
	return tokenString, nil
}