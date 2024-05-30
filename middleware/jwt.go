package middleware

import (
	"database/sql"
	"errors"
	"fmt"
	"orquideapp/src/domain"
	beneficiary_repository "orquideapp/src/interfaces/repository/beneficiary"
	employee_repository "orquideapp/src/interfaces/repository/employee"
	user_repository "orquideapp/src/interfaces/repository/user"
	beneficiary_usecase "orquideapp/src/usecase/beneficiary"
	employee_usecase "orquideapp/src/usecase/employee"
	user_usecase "orquideapp/src/usecase/user"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtMiddleware(fn func(c echo.Context, claims *domain.JwtCustomClaims) (bool, error)) echo.MiddlewareFunc {
	return middleware.KeyAuth(func(auth string, c echo.Context) (bool, error) {
		claims := &domain.JwtCustomClaims{}
		token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_JWT")), nil
		})
		if err != nil || !token.Valid {
			return false, err
		}

		return fn(c, claims)
	})
}

func JwtUserMiddleware(db *sql.DB) echo.MiddlewareFunc {
	repository := &user_repository.Repository{DB: db}
	uc := &user_usecase.UseCase{Repository: repository}
	return JwtMiddleware(func(c echo.Context, claims *domain.JwtCustomClaims) (bool, error) {
		user, err := uc.GetByEmail(claims.Email)
		if err != nil {
			return false, errors.New("error user")
		}
		if user.ID != claims.Id {
			return false, errors.New("error user")
		}
		c.Set("user", user)
		return true, nil
	})
}

func JwtEmployeMiddleware(db *sql.DB) echo.MiddlewareFunc {
	repository := &employee_repository.Repository{DB: db}
	uc := &employee_usecase.UseCase{Repository: repository}
	return JwtMiddleware(func(c echo.Context, claims *domain.JwtCustomClaims) (bool, error) {
		employee, err := uc.GetByEmail(claims.Email)
		if err != nil {
			return false, errors.New("error user")
		}
		if employee.ID != claims.Id {
			return false, errors.New("error user")
		}
		c.Set("employee", employee)
		return true, nil
	})
}

func JwtBeneficiaryMiddleware(db *sql.DB) echo.MiddlewareFunc {
	repository := &beneficiary_repository.Repository{DB: db}
	uc := &beneficiary_usecase.UseCase{Repository: repository}
	return JwtMiddleware(func(c echo.Context, claims *domain.JwtCustomClaims) (bool, error) {
		beneficiary, err := uc.GetByEmail(claims.Email)
		if err != nil {
			return false, errors.New("error user")
		}
		if beneficiary.ID != claims.Id {
			return false, errors.New("error user")
		}
		c.Set("beneficiary", beneficiary)
		return true, nil
	})
}
