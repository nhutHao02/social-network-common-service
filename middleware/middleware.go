package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	resErr "github.com/nhutHao02/social-network-common-service/utils/error"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/utils/token"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewUnauthorizedErrorResponse())
			return
		}

		claims, err := token.VerifyToken(tokenString)
		if err != nil {
			logger.Error("Token invalid ", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewUnauthorizedErrorResponse())
			return
		}

		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			logger.Error("Token Expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewUnauthorizedErrorResponse())
			return
		}

		c.Next()
	}
}

// Unary JWT interceptor for gRPC
func JWTUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Extract token from metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, resErr.NewResError(nil, "missing metadata")
	}

	// Extract the Authorization header
	authHeader, ok := md["Authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, resErr.NewResError(nil, "Authorization token not provided")
	}

	// Check if the token has the "Bearer " prefix
	tokenString := authHeader[0]
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return nil, resErr.NewResError(nil, "authorization token must have 'Bearer ' prefix")
	}

	// Remove "Bearer " prefix
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse and validate the token
	claims, err := token.VerifyToken(tokenString)
	if err != nil {
		logger.Error("Token invalid ", zap.Error(err))
		return nil, resErr.NewResError(nil, "Token invalid")
	}

	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		logger.Error("Token Expired")
		return nil, resErr.NewResError(nil, "Token Expired")
	}

	// Call the handler if the token is valid
	return handler(ctx, req)
}
