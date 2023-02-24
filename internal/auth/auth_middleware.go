package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"gin-starter/cmd/api/config"
	"gin-starter/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("Authorization header is missing")
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Authorization token is required"})
			return
		}

		var accessToken string

		tokenSplit := strings.Split(authHeader, "Bearer ")
		if len(tokenSplit) == 2 {
			accessToken = strings.TrimSpace(tokenSplit[1])
		} else {
			fmt.Println("Incorrect format of auth header")
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token format"})
			return
		}

		// TODO: implement blacklist for logout
		// if IsBlacklisted(accessToken) {
		// 	fmt.Println("Found in Blacklist")
		// 	utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token"})
		// 	return
		// }

		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JwtAccessTokenSecret), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("Invalid Token Signature")
				utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token signature"})
				return
			}
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token"})
			return
		}

		if !parsedToken.Valid {
			fmt.Println("Invalid Token")
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token"})
			return
		}

		// TODO: c.Set("user", user) huh?

		c.Next()
	}
}
