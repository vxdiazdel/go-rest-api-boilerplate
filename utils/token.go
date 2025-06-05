package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
)

type TokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func SignToken(userID uuid.UUID) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	claims := TokenClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signingKey)
}

func VerifyToken(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("verify token: %w", err)
	}
	if claims, ok := token.Claims.(*TokenClaims); ok {
		// check if token is expired
		exp, err := claims.GetExpirationTime()
		if err != nil {
			return nil, fmt.Errorf("get claims expiration time: %w", err)
		}
		if exp.Before(time.Now()) {
			return nil, ErrTokenExpired
		}

		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}
