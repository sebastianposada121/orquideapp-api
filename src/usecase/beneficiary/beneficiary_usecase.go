package beneficiary_usecase

import (
	"errors"
	"fmt"
	"orquideapp/src/domain"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository                domain.BeneficiaryRepository
	AppointmentStepRepository domain.AppointmentStepRepository
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

func (uc *UseCase) CreateMedicalAppointment(medicalAppointment domain.MedicalAppointment) error {
	medicalAppointments, err := uc.Repository.MedicalAppointments(medicalAppointment.BeneficiaryID)

	if err != nil {
		return err
	}

	for index := range medicalAppointments {

		if medicalAppointments[index].Active {
			fmt.Println(medicalAppointments[index])
			return errors.New("active medical appointment")
		}
	}

	id, err := uc.Repository.CreateMedicalAppointment(medicalAppointment)

	fmt.Println(id)
	if err != nil {
		return err
	}

	// appointmentStep := domain.AppointmentStep{
	// 	Name:                 "cita consultorio general",
	// 	MedicalAppointmentID: id,
	// 	Status:               "complete",
	// }

	// if err := uc.AppointmentStepRepository.Create(appointmentStep); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	fmt.Println(id)

	return err
}

func (uc *UseCase) GetByEmail(email string) (domain.Beneficiary, error) {
	return uc.Repository.GetByEmail(email)
}

func (uc UseCase) CreateAppointmentStepRepository(appointmentStep domain.AppointmentStep) error {

	return uc.AppointmentStepRepository.Create(appointmentStep)

}
