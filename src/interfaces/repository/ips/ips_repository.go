package ips_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetByID(id int) (domain.Ips, error) {
	var ips domain.Ips
	// query := "SELECT id, name, email, password FROM users WHERE id = ?"
	// row := r.DB.QueryRow(query, id)
	// err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return ips, nil
}

func (r *Repository) GetAllByBranchOfficeID(branchOfficeId int, id int) ([]domain.Ips, error) {
	ipsList := []domain.Ips{}
	query := `SELECT
    		hpi.id, hpi.name, hpi.email, hpi.phone, hpi.address, hpi.city_id, hpi.active
		FROM
    		users u
		INNER JOIN
    		health_service_branches hsb ON u.id = hsb.user_id
		INNER JOIN
    		health_provider_institutions hpi ON hsb.id = hpi.eps_id
		WHERE
    		u.id = ? AND hsb.id = ?;
		`
	data, err := r.DB.Query(query, branchOfficeId, id)
	if err != nil {
		return ipsList, err
	}

	for data.Next() {
		ips := domain.Ips{}
		data.Scan(&ips.ID, &ips.Name, &ips.Email, &ips.Phone, &ips.Address, &ips.CityId, &ips.Active)
		ipsList = append(ipsList, ips)
	}
	return ipsList, nil
}

func (r *Repository) Create(ips domain.Ips) error {
	query := "INSERT INTO health_provider_institutions (name, email, phone, address, city_id, type_id, eps_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, ips.Name, ips.Email, ips.Phone, ips.Address, ips.CityId, ips.TypeId, ips.EpsId)
	return err
}

func (r *Repository) Update(ips domain.Ips) error {
	// query := "UPDATE ipss SET name = ?, email = ?, password = ? WHERE id = ?"
	// _, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return nil
}

func (r *Repository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
