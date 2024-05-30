package beneficiary_usecase

import (
	"errors"
	"orquideapp/src/domain"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository                   domain.BeneficiaryRepository
	MedicalAppointmentRepository domain.MedicalAppointmentRepository
}

func (uc *UseCase) Create(beneficiary domain.Beneficiary) error {
	return uc.Repository.Create(beneficiary)
}

func (uc *UseCase) GetAllByBranchOfficeID(id int) (domain.Beneficiaries, error) {
	return uc.Repository.GetAllByBranchOfficeID(id)
}

func (uc *UseCase) UpdatePassword(credentials domain.Login) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 8)
	if err != nil {
		return err
	}

	credentials.Password = string(hashPassword)
	return uc.Repository.UpdatePassword(credentials)
}

func (uc *UseCase) GetAllMedicalAppointmentById(id int) (domain.MedicalAppointments, error) {
	return uc.MedicalAppointmentRepository.GetAllMedicalAppointmentByBeneficiaryID(id)
}

func (uc *UseCase) CreateMedicalAppointment(medicalAppointment domain.MedicalAppointment) error {
	medicalAppointments, err := uc.MedicalAppointmentRepository.GetAllMedicalAppointmentByBeneficiaryID(medicalAppointment.BeneficiaryID)

	if err != nil {
		return err
	}

	for index := range medicalAppointments {

		if medicalAppointments[index].Active {
			return errors.New("active medical appointment")
		}
	}

	id, err := uc.MedicalAppointmentRepository.CreateMedicalAppointment(medicalAppointment)

	if err != nil {
		return err
	}

	appointmentStep := domain.AppointmentStep{
		Name:                 "Confirmacion cita medica",
		MedicalAppointmentID: id,
		Status:               "pending",
	}

	if err := uc.MedicalAppointmentRepository.CreateAppointmentStep(appointmentStep); err != nil {
		return err
	}

	return err
}

func (uc *UseCase) GetByEmail(email string) (domain.Beneficiary, error) {
	return uc.Repository.GetByEmail(email)
}
