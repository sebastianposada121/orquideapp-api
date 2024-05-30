package beneficiary_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Create(beneficiary domain.Beneficiary) error {
	query := "INSERT INTO beneficiaries (name, last_name, document_id, document_number, email, phone, address, birthdate, gender_id, eps_id, city_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, beneficiary.Name, beneficiary.LastName, beneficiary.DocumentId, beneficiary.DocumentNumber, beneficiary.Email, beneficiary.Phone, beneficiary.Address, beneficiary.Birthdate, beneficiary.GenderId, beneficiary.EpsId, beneficiary.CityId)
	return err
}

func (r *Repository) GetAllByBranchOfficeID(id int) (domain.Beneficiaries, error) {

	beneficiaries := domain.Beneficiaries{}
	query := "SELECT id, name, email, phone, document_id, document_number FROM beneficiaries WHERE eps_id = ?"
	data, err := r.DB.Query(query, id)
	if err != nil {
		return beneficiaries, err
	}

	for data.Next() {
		beneficiary := domain.Beneficiary{}
		data.Scan(&beneficiary.ID, &beneficiary.Name, &beneficiary.Email, &beneficiary.Phone, &beneficiary.DocumentId, &beneficiary.DocumentNumber)
		beneficiaries = append(beneficiaries, beneficiary)
	}

	return beneficiaries, nil
}

func (r *Repository) UpdatePassword(credentials domain.Login) error {
	query := "UPDATE beneficiaries SET password = ? where email = ?"
	_, err := r.DB.Exec(query, credentials.Password, credentials.Email)
	return err
}

func (r *Repository) CreateMedicalAppointment(medicalEndpoint domain.MedicalAppointment) (int, error) {
	query := "INSERT INTO medical_appointments (beneficiary_id, appointment_date, service_id, description) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, medicalEndpoint.BeneficiaryID, medicalEndpoint.AppointmentDate, medicalEndpoint.ServiceID, medicalEndpoint.Description)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return int(id), err
	}

	return int(id), err
}

func (r *Repository) GetByEmail(email string) (domain.Beneficiary, error) {
	benefeciary := domain.Beneficiary{}
	query := "SELECT id, email, name, password FROM beneficiaries WHERE email = ?"
	row := r.DB.QueryRow(query, email)

	if err := row.Scan(&benefeciary.ID, &benefeciary.Email, &benefeciary.Name, &benefeciary.Password); err != nil {
		return benefeciary, err
	}

	return benefeciary, nil
}

func (r *Repository) MedicalAppointments(id int) (domain.MedicalAppointments, error) {
	medicalAppointmens := domain.MedicalAppointments{}
	query := "SELECT id, active, description, service_id, appointment_date FROM medical_appointments WHERE beneficiary_id = ?"
	data, err := r.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	for data.NextResultSet() {
		medicalAppointment := domain.MedicalAppointment{}
		data.Scan(&medicalAppointment.ID, &medicalAppointment.Active, &medicalAppointment.Description, &medicalAppointment.ServiceID, &medicalAppointment.AppointmentDate)
		medicalAppointmens = append(medicalAppointmens, medicalAppointment)
	}

	return medicalAppointmens, nil
}
