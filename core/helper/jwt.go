package helper

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = ConfigJson()["development"].(map[string]interface{})["key"]

func GenerateToken(claims *jwt.MapClaims) string{
	key := fmt.Sprint(secretKey)
	keyBytes := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString(keyBytes)
	if err != nil {
		panic("helper - GenerateToken, token.SignedString: " + err.Error())
	}

	return webToken
}

func VerfyToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		key := fmt.Sprint(secretKey)

		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerfyToken(tokenString)
	if err != nil {
		panic("helper - DecodeToken, VerfyToken: " + err.Error())
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}