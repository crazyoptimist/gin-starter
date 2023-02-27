package auth

import (
	"fmt"
	"gin-starter/cmd/api/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessTokenKeyId = iota + 1
	RefreshTokenKeyId
)

func GenerateAccessToken(userId uint) (string, error) {

	secretKey := []byte(config.Config.JwtAccessTokenSecret)
	expiresIn := config.Config.JwtAccessTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = AccessTokenKeyId

	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(expiresIn)
	fmt.Println(claims)
	claims["iss"] = userId

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateRefreshToken(userId uint) (string, error) {

	secretKey := []byte(config.Config.JwtRefreshTokenSecret)
	expiresIn := config.Config.JwtRefreshTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = RefreshTokenKeyId

	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(expiresIn)
	claims["iss"] = userId

	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func GenerateTokenPair(userId uint) (string, string, error) {

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

func ValidateToken(tokenString string) (isValid bool, userId uint, keyId uint, err error) {

	var key []byte

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		keyId = token.Header["kid"].(uint)

		switch keyId {
		case AccessTokenKeyId:
			key = []byte(config.Config.JwtAccessTokenSecret)
		case RefreshTokenKeyId:
			key = []byte(config.Config.JwtRefreshTokenSecret)
		}

		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Invalid Token Signature")
			return
		}
		return
	}

	if !token.Valid {
		fmt.Println("Invalid Token")
		return
	}

	isValid = true
	userId = claims["iss"].(uint)
	return
}
