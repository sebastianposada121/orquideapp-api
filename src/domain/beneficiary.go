package domain

import (
	"time"
)

type BeneficiaryRepository interface {
	GetByEmail(email string) (Beneficiary, error)
	Create(user Beneficiary) error
	GetAllByBranchOfficeID(branchOfficeId int) (Beneficiaries, error)
	UpdatePassword(credentials Login) error
}

type Beneficiary struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	DocumentId     int       `json:"document_id"`
	DocumentNumber string    `json:"document_number"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	GenderId       int       `json:"gender_id"`
	Birthdate      time.Time `json:"birthdate"`
	RegisterDate   time.Time `json:"register_date"`
	EpsId          int       `json:"eps_id"`
	CityId         int       `json:"city_id"`
}

type Beneficiaries []Beneficiary
