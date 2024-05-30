package branchOffice_usecase

import (
	"errors"
	"orquideapp/src/domain"
)

type UseCase struct {
	Repository domain.BranchOfficeRepository
}

func (uc *UseCase) Create(branchOffice domain.BranchOffice) error {

	if _, error := uc.Repository.GetByEmail(branchOffice.Email); error == nil {
		return errors.New("branch office already exists")
	}
	return uc.Repository.Create(branchOffice)
}

func (uc *UseCase) GetAllByUserId(id int) (domain.BranchOffices, error) {
	return uc.Repository.GetAllByUserId(id)
}
