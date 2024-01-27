package middleware

import (
	"net/http"
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

		// TODO: Implement token blacklist for logout, using Redis or another in-memory db?
		// if helper.IsBlacklistedJWT(accessToken) {
		// 	utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid authorization token"})
		// 	return
		// }

		claims := jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JwtAccessTokenSecret), nil
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

		// TODO: Inject user object to the context?
		// c.Set("user", user)

		c.Next()
	}
}
