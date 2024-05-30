package user_usecase

import (
	"orquideapp/src/domain"
)

type UseCase struct {
	Repository domain.UserRepository
}

func (uc *UseCase) GetByEmail(email string) (domain.User, error) {
	return uc.Repository.GetByEmail(email)
}
