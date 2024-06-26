package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-starter/internal/config"
	"gin-starter/internal/domain/auth"
	"gin-starter/internal/domain/model"
	"gin-starter/internal/infrastructure/controller"
	"gin-starter/pkg/common"
)

const CACHE_USER_TTL = 24 * time.Hour

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			raiseUnauthorizedError(c, "Authorization header is required")
			return
		}

		var accessToken string

		tokenSplit := strings.Split(authHeader, "Bearer ")
		if len(tokenSplit) == 2 {
			accessToken = strings.TrimSpace(tokenSplit[1])
		} else {
			raiseUnauthorizedError(c, "Invalid authorization header")
			return
		}

		isJwtValid, userIdString, _, err := auth.ValidateJwtToken(accessToken)
		if err != nil {
			common.RaiseHttpError(
				c,
				http.StatusUnauthorized,
				err,
			)
			return
		}
		if !isJwtValid {
			raiseUnauthorizedError(c, "Invalid access token")
			return
		}

		var user model.User

		// Check if there's a cached user
		userJson, err := config.Global.RedisClient.Get(
			c,
			controller.CACHE_KEY_PREFIX_USER+userIdString,
		).Result()
		if err != nil {
			// If there's no cached user, query DB
			if err := config.Global.DB.Where(
				"id = ?", userIdString,
			).First(&user).Error; err != nil {
				raiseUnauthorizedError(c, "User not found")
				return
			}
			// And cache the user
			userJson, err := json.Marshal(user)
			if err != nil {
				raiseUnauthorizedError(c, "Marshaling user failed")
				return
			}
			if err := config.Global.RedisClient.Set(
				c,
				controller.CACHE_KEY_PREFIX_USER+userIdString,
				userJson,
				CACHE_USER_TTL,
			).Err(); err != nil {
				raiseUnauthorizedError(c, "Caching user failed")
				return
			}
		} else {
			// If a cached user exists, use it
			err = json.Unmarshal([]byte(userJson), &user)
			if err != nil {
				raiseUnauthorizedError(c, "Unmarshaling cached user failed")
			}
		}

		// Include the user to the context
		c.Set("user", user)

		c.Next()
	}
}

func raiseUnauthorizedError(c *gin.Context, message string) {
	common.RaiseHttpError(
		c,
		http.StatusUnauthorized,
		errors.New(message),
	)
}
