package branchOffice_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Create(branchOffice domain.BranchOffice) error {
	query := "INSERT INTO health_service_branches (name, email, city_id, phone, address, user_id) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, branchOffice.Name, branchOffice.Email, branchOffice.CityId, branchOffice.Phone, branchOffice.Address, branchOffice.UserId)
	return err
}

func (r *Repository) Update(branchOffice domain.BranchOffice) error {
	// query := "UPDATE ipss SET name = ?, email = ?, password = ? WHERE id = ?"
	// _, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return nil
}

func (r *Repository) GetByEmail(email string) (domain.BranchOffice, error) {
	var branchOffice domain.BranchOffice
	query := "SELECT id, name, email, password FROM health_service_branches WHERE email = ?"
	row := r.DB.QueryRow(query, email)
	err := row.Scan(&branchOffice.ID, &branchOffice.Name, &branchOffice.Email)
	return branchOffice, err
}

func (r *Repository) GetAllByUserId(id int) (domain.BranchOffices, error) {

	branchOffices := domain.BranchOffices{}
	query := "SELECT id, name, email, phone, address, city_id FROM health_service_branches WHERE user_id = ?"
	data, err := r.DB.Query(query, id)
	if err != nil {
		return branchOffices, err
	}

	for data.Next() {
		branchOffice := domain.BranchOffice{}
		data.Scan(&branchOffice.ID, &branchOffice.Name, &branchOffice.Email, &branchOffice.Phone, &branchOffice.Address, &branchOffice.CityId)
		branchOffices = append(branchOffices, branchOffice)
	}

	return branchOffices, nil
}
