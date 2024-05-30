package domain

import "time"

type AppointmentStepRepository interface {
	Create(appointmentStep AppointmentStep) error
	Update(appointmentStep AppointmentStep) error
	Delete(id int) error
	GetAllByMedicalAppointmentID(id int) (AppointmentSteps, error)
}

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
