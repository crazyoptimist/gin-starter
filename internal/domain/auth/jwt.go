package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gin-starter/internal/config"
	"gin-starter/pkg/common"
)

const (
	AccessTokenKeyId int = iota + 1
	RefreshTokenKeyId
)

func GenerateJwtToken(keyId int, userId int) (string, error) {
	var secretKey []byte
	var expiresIn time.Duration

	switch keyId {
	case RefreshTokenKeyId:
		secretKey = []byte(config.Global.JwtRefreshTokenSecret)
		expiresIn = config.Global.JwtRefreshTokenExpiresIn
	case AccessTokenKeyId:
		secretKey = []byte(config.Global.JwtAccessTokenSecret)
		expiresIn = config.Global.JwtAccessTokenExpiresIn
	default:
		keyId = AccessTokenKeyId
		secretKey = []byte(config.Global.JwtAccessTokenSecret)
		expiresIn = config.Global.JwtAccessTokenExpiresIn
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = keyId

	claims := token.Claims.(jwt.MapClaims)
	issuedAt := time.Now()
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = issuedAt.Add(expiresIn).Unix()
	claims["sub"] = strconv.Itoa(int(userId))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenPair(userId int) (accessToken, refreshToken string, err error) {

	accessToken, err = GenerateJwtToken(AccessTokenKeyId, userId)
	if err != nil {
		return
	}

	refreshToken, err = GenerateJwtToken(RefreshTokenKeyId, userId)
	if err != nil {
		return
	}

	return
}

func ValidateToken(tokenString string) (isValid bool, userId int, keyId int, err error) {

	var key []byte

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			keyId = int(token.Header["kid"].(float64))

			switch keyId {
			case AccessTokenKeyId:
				key = []byte(config.Global.JwtAccessTokenSecret)
			case RefreshTokenKeyId:
				key = []byte(config.Global.JwtRefreshTokenSecret)
			}

			return key, nil
		},
	)

	if err != nil {
		common.Logger.Error("JWT validation failed: ", err)
		return
	}

	if !token.Valid {
		err = errors.New("Invalid JWT token")
		return
	}

	isValid = true

	sub, err := claims.GetSubject()
	userId, err = strconv.Atoi(sub)

	return
}
