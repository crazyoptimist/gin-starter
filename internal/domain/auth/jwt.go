package auth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gin-starter/internal/config"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

const (
	AccessTokenKeyId uint = iota + 1
	RefreshTokenKeyId
)

func GenerateAccessToken(userId uint) (string, error) {

	secretKey := []byte(config.Global.JwtAccessTokenSecret)
	expiresIn := config.Global.JwtAccessTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = AccessTokenKeyId

	claims := token.Claims.(jwt.MapClaims)
	issuedAt := time.Now()
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expiresIn).Unix()
	claims["iss"] = strconv.Itoa(int(userId))

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func GenerateRefreshToken(userId uint) (string, error) {

	secretKey := []byte(config.Global.JwtRefreshTokenSecret)
	expiresIn := config.Global.JwtRefreshTokenExpiresIn

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = RefreshTokenKeyId

	claims := token.Claims.(jwt.MapClaims)
	issuedAt := time.Now()
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expiresIn).Unix()
	claims["iss"] = strconv.Itoa(int(userId))

	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func GenerateTokenPair(userId uint) (accessToken, refreshToken string, err error) {

	accessToken, err = GenerateAccessToken(userId)
	if err != nil {
		return
	}

	refreshToken, err = GenerateRefreshToken(userId)
	if err != nil {
		return
	}

	return
}

func ValidateToken(tokenString string) (isValid bool, userId uint, keyId uint, err error) {

	var key []byte

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			// TODO: refactor
			keyId = uint(token.Header["kid"].(float64))

			switch keyId {
			case AccessTokenKeyId:
				key = []byte(config.Global.JwtAccessTokenSecret)
			case RefreshTokenKeyId:
				key = []byte(config.Global.JwtRefreshTokenSecret)
			}

			return key, nil
		})

	if err != nil {
		common.Logger.Error("JWT validation failed: ", err)
		return
	}

	if !token.Valid {
		err = &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid Token"}
		return
	}

	isValid = true

	issuer, _ := claims.GetIssuer()
	issuerAsInt, _ := strconv.Atoi(issuer)
	userId = uint(issuerAsInt)

	return
}
