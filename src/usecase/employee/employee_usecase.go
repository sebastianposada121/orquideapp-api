package employee_usecase

import (
	"orquideapp/src/domain"
	"orquideapp/utils"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository                   domain.EmployeesRepository
	MedicalAppointmentRepository domain.MedicalAppointmentRepository
}

func (uc *UseCase) Create(employee domain.Employee) error {
	return uc.Repository.Create(employee)
}

func (uc *UseCase) CreateRol(rol domain.Rol) error {
	return uc.Repository.CreateRol(rol)
}

func (uc *UseCase) GetByEmail(email string) (domain.Employee, error) {
	return uc.Repository.GetByEmail(email)
}

func (uc *UseCase) Login(credentials domain.Login) (string, error) {
	employee, err := uc.Repository.GetByEmail(credentials.Email)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(credentials.Password)); err != nil {
		return "", err
	}

	token, err := utils.GenerateJwt(employee.ID, employee.Email, employee.Name, employee.IpsId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UseCase) UpdatePassword(credentials domain.Login) error {
	return uc.Repository.UpdatePassword(credentials)
}

func (uc *UseCase) CreateAppointmentStep(appointmentStep domain.AppointmentStep) error {
	return uc.MedicalAppointmentRepository.CreateAppointmentStep(appointmentStep)
}
