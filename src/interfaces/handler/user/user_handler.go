package user_handler

import (
	"net/http"
	user_usecase "orquideapp/src/usecase/user"
	"orquideapp/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase *user_usecase.UseCase
}

func (h *Handler) GetByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := h.UseCase.GetByEmail(email)
	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "No found", nil)
	}
	return utils.GenericResponse(c, http.StatusOK, true, "Success user", user)
}
