package token

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("nhuHao02-socialNetwork")

func CreateToken(idUser string) (string, error) {
	// create token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  idUser,
		"iss": "nhutHao02/socialNetwork",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	// signature with secretKey
	tokenString, err := claims.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	return claims, nil
}

func GetUserId(c *gin.Context) (int, error) {
	// get jwt token
	token, err := GetTokenString(c)
	if err != nil {
		return 0, errors.New("TOKEN_IS_MISSING")
	}

	claims, err := VerifyToken(token)
	if err != nil {
		return 0, err
	}
	idStr, ok := claims["id"].(string)
	if !ok {
		return 0, errors.New("INVALID_TOKEN_CLAIM_TYPE")
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("INVALID_USER_ID")
	}

	return idInt, nil
}

func GetTokenString(c *gin.Context) (string, error) {
	tokenHeaderName := "Bearer "
	authHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authHeader, tokenHeaderName) {
		return "", errors.New("DOES_NOT_EXIST_Bearer")
	}
	tokenString := authHeader[len(tokenHeaderName):]

	return tokenString, nil
}
