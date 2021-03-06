package middlewares

import (
	"net/http"
	"strings"

	"github.com/fetch_app/constants"
	responses "github.com/fetch_app/handlers/response"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware return gin middleware for authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ParseToken(ctx)
		if err != nil {
			unauthorizedResponse := responses.Response{
				Errors:  []string{err.Error()},
				Message: "Unauthorized data",
				Status:  constants.StatusFailed,
			}

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		ctx.Set(constants.JwtAppKey, claims)
		ctx.Next()
	}
}

// ParseToken :nodoc:
func ParseToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		return nil, ginJwt.ErrEmptyAuthHeader
	}

	parts := strings.SplitN(tokenString, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, ginJwt.ErrInvalidAuthHeader
	}

	tokenString = parts[1]
	jwtToken, err := jwt.Parse(tokenString, tokenParser)
	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

func tokenParser(token *jwt.Token) (interface{}, error) {
	if jwt.GetSigningMethod(constants.AuthSigningMethod) != token.Method {
		return nil, ginJwt.ErrInvalidSigningAlgorithm
	}

	return []byte(constants.JwtAppKey), nil
}
