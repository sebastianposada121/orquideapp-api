package ips_usecase

import (
	"orquideapp/src/domain"
)

type UseCase struct {
	Repository domain.IpsRepository
}

func (uc *UseCase) Create(ips domain.Ips) error {
	return uc.Repository.Create(ips)
}

func (uc *UseCase) GetAllByBranchOfficeID(userId int, id int) ([]domain.Ips, error) {
	return uc.Repository.GetAllByBranchOfficeID(userId, id)
}

func (uc *UseCase) GetByID(id int) (domain.Ips, error) {
	return uc.Repository.GetByID(id)
}
