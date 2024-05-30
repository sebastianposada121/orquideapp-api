package domain

import "time"

type MedicalAppointmentRepository interface {
	// medical appointments
	CreateMedicalAppointment(medicalAppointment MedicalAppointment) (int, error)
	GetAllMedicalAppointmentByBeneficiaryID(id int) (MedicalAppointments, error)

	// steps
	CreateAppointmentStep(appointmentStep AppointmentStep) error
	GetAllAppointmentStepsByMedicalAppointmentID(id int) (AppointmentSteps, error)
}

type MedicalAppointment struct {
	ID              int       `json:"id"`
	BeneficiaryID   int       `json:"beneficiary_id"`
	Active          bool      `json:"active"`
	ServiceID       int       `json:"service_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	Description     string    `json:"description"`
}

type MedicalAppointments []MedicalAppointment

type AppointmentStep struct {
	ID                        int       `json:"id"`
	MedicalAppointmentID      int       `json:"medical_appointment_id"`
	ChildMedicalAppointmentId int       `json:"child_medical_appointment_id"`
	Name                      string    `json:"name"`
	Description               string    `json:"description"`
	Status                    string    `json:"status"`
	CreatedAt                 time.Time `json:"create_at"`
}

type AppointmentSteps []AppointmentStep
