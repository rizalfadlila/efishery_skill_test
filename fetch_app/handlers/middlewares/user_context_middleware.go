package middlewares

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/fetch_app/constants"
	"github.com/gin-gonic/gin"
)

// UserContextMiddleware extracts user context from JWT and injects it to request context
func UserContextMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		authHeader := ginCtx.GetHeader("Authorization")
		if len(authHeader) < 7 {
			ginCtx.Next()
			return
		}

		var claims jwt.MapClaims
		_, _ = jwt.ParseWithClaims(authHeader[7:], &claims, nil)

		requestCtx := ginCtx.Request.Context()

		if value, ok := claims["name"].(string); ok {
			requestCtx = context.WithValue(requestCtx, constants.KeyName, value)
		}

		if value, ok := claims["phone"].(string); ok {
			requestCtx = context.WithValue(requestCtx, constants.KeyPhone, value)
		}

		if value, ok := claims["role"].(string); ok {
			requestCtx = context.WithValue(requestCtx, constants.KeyRole, value)
		}

		if value, ok := claims["timestamp"].(string); ok {
			requestCtx = context.WithValue(requestCtx, constants.KeyTimestamp, value)
		}

		ginCtx.Request = ginCtx.Request.WithContext(requestCtx)
		ginCtx.Next()
	}
}
