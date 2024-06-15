package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"gin-starter/internal/config"
	"gin-starter/internal/domain/model"
	"gin-starter/internal/infrastructure/controller"
	"gin-starter/pkg/utils"
)

const CACHE_USER_TTL = 24 * time.Hour

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

		issuer, err := parsedToken.Claims.GetIssuer()
		if err != nil {
			utils.RaiseHttpError(c, http.StatusUnauthorized, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Getting issuer from the parsed token failed"})
			return

		}

		var user model.User

		// Check if there's a cached user
		userJson, err := config.Global.RedisClient.Get(
			c,
			controller.CACHE_KEY_PREFIX_USER+issuer,
		).Result()
		if err != nil {
			// If there's no cached user, query DB
			if err := config.Global.DB.Where(
				"id = ?", issuer,
			).First(&user).Error; err != nil {
				utils.RaiseHttpError(
					c,
					http.StatusUnauthorized,
					&utils.HttpError{
						Code:    http.StatusUnauthorized,
						Message: "User not found",
					},
				)
				return
			}
			// And cache the user
			userJson, err := json.Marshal(user)
			if err != nil {
				utils.RaiseHttpError(
					c,
					http.StatusUnauthorized,
					&utils.HttpError{
						Code:    http.StatusUnauthorized,
						Message: "Marshaling user failed",
					},
				)
				return
			}
			if err := config.Global.RedisClient.Set(
				c,
				controller.CACHE_KEY_PREFIX_USER+issuer,
				userJson,
				CACHE_USER_TTL,
			).Err(); err != nil {
				utils.RaiseHttpError(
					c,
					http.StatusUnauthorized,
					&utils.HttpError{
						Code:    http.StatusUnauthorized,
						Message: "Caching user failed",
					},
				)
				return
			}
		} else {
			// If a cached user exists, use it
			err = json.Unmarshal([]byte(userJson), &user)
			if err != nil {
				utils.RaiseHttpError(
					c,
					http.StatusUnauthorized,
					&utils.HttpError{
						Code:    http.StatusUnauthorized,
						Message: "Unmarshaling cached user failed",
					},
				)
			}
		}

		// Include the user to the context
		c.Set("user", user)

		c.Next()
	}
}
