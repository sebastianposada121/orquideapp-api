package medicalAppointment_usecase

import "orquideapp/src/domain"

type UseCase struct {
	Repository domain.MedicalAppointmentRepository
}

func (uc UseCase) CreateMedicalAppointment(medicalEndpoint domain.MedicalAppointment) (int, error) {
	return uc.Repository.CreateMedicalAppointment(medicalEndpoint)
}

func (uc UseCase) GetAllMedicalAppointmentByBeneficiaryID(id int) (domain.MedicalAppointments, error) {
	return uc.Repository.GetAllMedicalAppointmentByBeneficiaryID(id)
}

func (uc UseCase) CreateAppointmentStep(appointmentStep domain.AppointmentStep) error {
	return uc.Repository.CreateAppointmentStep(appointmentStep)
}

func (uc UseCase) GetAllAppointmentStepsByMedicalAppointmentID(id int) (domain.AppointmentSteps, error) {
	return uc.Repository.GetAllAppointmentStepsByMedicalAppointmentID(id)
}
