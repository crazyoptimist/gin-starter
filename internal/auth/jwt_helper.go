package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gin-starter/internal/core/config"
	"gin-starter/internal/core/logger"
	"gin-starter/pkg/utils"
)

const (
	AccessTokenKeyId uint = iota + 1
	RefreshTokenKeyId
)

func GenerateAccessToken(userId uint) (string, error) {

	secretKey := []byte(config.Config.JwtAccessTokenSecret)
	expiresIn := config.Config.JwtAccessTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = AccessTokenKeyId

	claims := token.Claims.(jwt.MapClaims)
	issuedAt := time.Now()
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expiresIn).Unix()
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
	issuedAt := time.Now()
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expiresIn).Unix()
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

		keyId = uint(token.Header["kid"].(float64))

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
			logger.Logger.Info(err)
			return
		}
		logger.Logger.Info(err)
		return
	}

	if !token.Valid {
		err = &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid Token"}
		logger.Logger.Info("Invalid Token")
		return
	}

	isValid = true

	userId = uint(claims["iss"].(float64))

	return
}
