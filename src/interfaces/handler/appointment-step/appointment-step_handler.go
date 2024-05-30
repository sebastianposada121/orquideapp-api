package appointmentstep_handler

import (
	"orquideapp/src/domain"
	appointmentstep_usecase "orquideapp/src/usecase/appointment-step"
)

type Handler struct {
	UseCase appointmentstep_usecase.UseCase
}

func (h Handler) Create(appointmentStep domain.AppointmentStep) error {
	return h.UseCase.Create(appointmentStep)
}

func (h Handler) Delete(id int) error {
	return h.UseCase.Delete(id)
}

func (h Handler) Update(appointmentStep domain.AppointmentStep) error {
	return h.UseCase.Update(appointmentStep)
}

func (h Handler) GetAllByMedicalAppointmentID(benefeciaryID int) (domain.AppointmentSteps, error) {
	return h.UseCase.GetAllByMedicalAppointmentID(benefeciaryID)
}
