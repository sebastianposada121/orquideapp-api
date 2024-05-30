package appointmentstep_usecase

import "orquideapp/src/domain"

type UseCase struct {
	Repository domain.AppointmentStepRepository
}

func (uc UseCase) Create(appointmentStep domain.AppointmentStep) error {
	return uc.Repository.Create(appointmentStep)
}

func (uc UseCase) Update(appointmentStep domain.AppointmentStep) error {
	return uc.Repository.Update(appointmentStep)
}

func (uc UseCase) Delete(id int) error {
	return uc.Repository.Delete(id)
}

func (uc UseCase) GetAllByMedicalAppointmentID(id int) (domain.AppointmentSteps, error) {
	return nil, nil
}
