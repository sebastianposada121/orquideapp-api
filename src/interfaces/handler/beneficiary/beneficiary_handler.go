package beneficiary_handler

import (
	"net/http"
	"orquideapp/src/domain"
	beneficiary_usecase "orquideapp/src/usecase/beneficiary"
	"orquideapp/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase *beneficiary_usecase.UseCase
}

func (h *Handler) Create(c echo.Context) error {
	beneficiary := new(domain.Beneficiary)

	if err := c.Bind(beneficiary); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error create", nil)
	}

	if err := h.UseCase.Create(*beneficiary); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error create", nil)
	}

	return utils.GenericResponse(c, http.StatusCreated, true, "success", nil)
}

func (h *Handler) GetAll(c echo.Context) error {

	id, err := strconv.Atoi(c.QueryParam("branchOfficeId"))

	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error id", nil)
	}

	beneficiaries, err := h.UseCase.GetAllByBranchOfficeID(id)
	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error get", nil)
	}

	return utils.GenericResponse(c, http.StatusAccepted, true, "success", beneficiaries)
}

func (h *Handler) CreateMedicalAppointment(c echo.Context) error {

	medicalAppointment := new(domain.MedicalAppointment)
	beneficiary := c.Get("beneficiary").(domain.Beneficiary)

	if err := c.Bind(medicalAppointment); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error createad medical appointment", nil)
	}

	medicalAppointment.BeneficiaryID = beneficiary.ID
	if err := h.UseCase.CreateMedicalAppointment(*medicalAppointment); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}

	return utils.GenericResponse(c, http.StatusCreated, true, "create medical appointment", nil)
}
