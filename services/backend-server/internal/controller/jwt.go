package controller

import (
	"crypto/rand"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const secretLength = 25

var Secret []byte

func init() {
	secret := make([]byte, secretLength)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}

	Secret = secret
}

type JWTCustomClaims struct {
	UserUUID uuid.UUID `json:"name"`
	jwt.RegisteredClaims
}

func NewClaim(userUUID uuid.UUID) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaims{
		UserUUID: userUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.ExpriationDuration)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func NewJWTMiddlewareFunc() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(JWTCustomClaims)
			},
			SigningKey: Secret,
		},
	)
}
