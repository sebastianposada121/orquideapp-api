package domain

import "time"

type BranchOfficeRepository interface {
	Create(branchOffice BranchOffice) error
	Update(branchOffice BranchOffice) error
	GetByEmail(email string) (BranchOffice, error)
	GetAllByUserId(id int) (BranchOffices, error)
}

type BranchOffice struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	RegisterDate time.Time `json:"register_date"`
	CityId       int       `json:"city_id"`
	UserId       int       `json:"user_id"`
}

type BranchOffices []BranchOffice
