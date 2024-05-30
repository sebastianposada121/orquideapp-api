package branchOffice_hander

import (
	"net/http"
	"orquideapp/src/domain"
	branchOffice_usecase "orquideapp/src/usecase/branch-office"

	"orquideapp/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase branchOffice_usecase.UseCase
}

func (h *Handler) Create(c echo.Context) error {
	user := c.Get("user").(domain.User)
	branchOffice := new(domain.BranchOffice)
	if err := c.Bind(branchOffice); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error create branch office", nil)
	}
	branchOffice.UserId = user.ID
	if err := h.UseCase.Create(*branchOffice); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error create branch office", nil)
	}
	return utils.GenericResponse(c, http.StatusCreated, true, "branch office create successfully", nil)
}

func (h *Handler) GetAllByUserId(c echo.Context) error {
	user := c.Get("user").(domain.User)
	branchOffices, err := h.UseCase.GetAllByUserId(user.ID)
	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error get branch office", nil)
	}
	return utils.GenericResponse(c, http.StatusOK, true, "Success get branch office", branchOffices)
}
