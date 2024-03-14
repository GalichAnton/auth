package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/GalichAnton/auth/internal/models/claims"
)

// GenerateToken ...
func GenerateToken(info claims.UserClaims, secretKey []byte, duration time.Duration) (string, error) {
	claim := claims.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Email: info.Email,
		Role:  info.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(secretKey)
}

// VerifyToken ...
func VerifyToken(tokenStr string, secretKey []byte) (*claims.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claim, ok := token.Claims.(*claims.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claim, nil
}
