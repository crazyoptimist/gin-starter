package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"gin-starter/internal/config"
	"gin-starter/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Authorization token is required"})
			return
		}

		var accessToken string

		tokenSplit := strings.Split(authHeader, "Bearer ")
		if len(tokenSplit) == 2 {
			accessToken = strings.TrimSpace(tokenSplit[1])
		} else {
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token format"})
			return
		}

		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Global.JwtAccessTokenSecret), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token signature"})
				return
			}
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token"})
			return
		}

		if !parsedToken.Valid {
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token"})
			return
		}

		issuer, _ := parsedToken.Claims.GetIssuer()
		userId, _ := strconv.Atoi(issuer)
		c.Set("user", userId)

		c.Next()
	}
}
