package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/utils/token"
	"go.uber.org/zap"
)

type UnauthorizedErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewUnauthorizedErrorResponse() *UnauthorizedErrorResponse {
	return &UnauthorizedErrorResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid token",
	}
}

func JwtAuthMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := token.GetTokenString(c)
		if err != nil {
			logger.Error("Token invalid", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewUnauthorizedErrorResponse)
			return
		}

		claims, err := token.VerifyToken(tokenString)
		if err != nil {
			logger.Error("Token invalid ", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewUnauthorizedErrorResponse)
			return
		}

		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			logger.Error("Token Expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewUnauthorizedErrorResponse)
			return
		}

		c.Next()
	}
}
