package auth

import (
	"fmt"
	"gin-starter/cmd/api/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(userId string) (string, error) {

	secretKey := config.Config.JwtAccessTokenSecret
	expiresIn := config.Config.JwtAccessTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = "access_token"

	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(expiresIn * time.Second)
	claims["iss"] = userId

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateRefreshToken(userId string) (string, error) {

	secretKey := config.Config.JwtRefreshTokenSecret
	expiresIn := config.Config.JwtRefreshTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = "refresh_token"

	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(expiresIn * time.Second)
	claims["iss"] = userId

	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func GenerateTokenPair(userId string) (string, string, error) {

	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateToken(tokenString string) (bool, string, string, error) {

	var key []byte

	var keyID string

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		keyID = token.Header["kid"].(string)

		switch keyID {
		case "access_token":
			key = []byte(config.Config.JwtAccessTokenSecret)
		case "refresh_token":
			key = []byte(config.Config.JwtRefreshTokenSecret)
		}

		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Invalid Token Signature")
			return false, "", keyID, err
		}
		return false, "", keyID, err
	}

	if !token.Valid {
		fmt.Println("Invalid Token")
		return false, "", keyID, err
	}

	return true, claims["iss"].(string), keyID, nil
}
