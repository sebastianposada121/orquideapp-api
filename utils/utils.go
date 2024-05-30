package utils

import (
	"orquideapp/src/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Successful bool
	Message    string
	Data       interface{}
}

func GenericResponse(c echo.Context, status int, successful bool, message string, data interface{}) error {
	return c.JSON(status, Response{
		Successful: successful,
		Message:    message,
		Data:       data,
	})
}

func GenerateJwt(id int, name string, email string, ipsId int) (string, error) {
	claims := &domain.JwtCustomClaims{
		Name:  name,
		Email: email,
		IpsID: ipsId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
