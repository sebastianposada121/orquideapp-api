package domain

import "time"

type IpsRepository interface {
	GetByID(id int) (Ips, error)
	Create(ips Ips) error
	Update(ips Ips) error
	Delete(id int) error
	GetAllByBranchOfficeID(BranchOfficeId int, id int) ([]Ips, error)
}

type Ips struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	CityId       int       `json:"city_id"`
	TypeId       int       `json:"type_id"`
	RegisterDate time.Time `json:"register_date"`
	Active       bool      `json:"active"`
	EpsId        int       `json:"eps_id"`
}
