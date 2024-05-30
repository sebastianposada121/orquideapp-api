package auth_handler

import (
	"net/http"
	"orquideapp/src/domain"
	auth_usecase "orquideapp/src/usecase/auth"
	"orquideapp/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	UseCase *auth_usecase.UseCase
}

func (h *Handler) UserLogin(c echo.Context) error {
	login := new(domain.Login)
	if err := c.Bind(&login); err != nil {
		return utils.GenericResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
	}
	token, err := h.UseCase.UserLogin(*login)
	if err != nil {
		return utils.GenericResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
	}
	return utils.GenericResponse(c, http.StatusAccepted, false, "Success login", token)
}

func (h *Handler) EmployeeLogin(c echo.Context) error {
	login := new(domain.Login)
	if err := c.Bind(&login); err != nil {
		return utils.GenericResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
	}
	token, err := h.UseCase.EmployeeLogin(*login)
	if err != nil {
		return utils.GenericResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
	}
	return utils.GenericResponse(c, http.StatusAccepted, false, "Success login", token)
}

func (h *Handler) BeneficiaryLogin(c echo.Context) error {
	login := new(domain.Login)
	if err := c.Bind(&login); err != nil {
		return utils.GenericResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
	}
	token, err := h.UseCase.BeneficiaryLogin(*login)
	if err != nil {
		return utils.GenericResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
	}
	return utils.GenericResponse(c, http.StatusAccepted, false, "Success login", token)
}

func (h *Handler) UserSignup(c echo.Context) error {
	user := new(domain.User)

	if err := c.Bind(user); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "Error signup user", nil)
	}

	if err := h.UseCase.UserSignup(*user); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "Error signup user", nil)
	}

	return utils.GenericResponse(c, http.StatusCreated, true, "User created successfully", nil)
}

func (h *Handler) UpdateBeneficiaryPassword(c echo.Context) error {
	return updatePassword(c, func(credentials domain.Login) error {
		return h.UseCase.UpdateBeneficiaryPassword(credentials)
	})
}

func (h *Handler) UpdateEmployeePassword(c echo.Context) error {
	return updatePassword(c, func(credentials domain.Login) error {
		return h.UseCase.UpdateEmployeePassword(credentials)
	})
}

func updatePassword(c echo.Context, update func(credentials domain.Login) error) error {
	credentials := new(domain.Login)

	token := c.QueryParam("token")

	if token != "secret" {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error update password", nil)
	}

	if err := c.Bind(credentials); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error update password", nil)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 8)
	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error update password", nil)
	}

	credentials.Password = string(hashPassword)

	if err := update(*credentials); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error update password", nil)
	}

	return utils.GenericResponse(c, http.StatusOK, true, "success update password", nil)
}
