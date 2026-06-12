package k9_health

type CreateK9HealthRecordRequest struct {

	K9ID string `json:"k9_id" binding:"required"`

	VisitedDate string `json:"visited_date" binding:"required"` 

	ReleasedDate string `json:"released_date"`

	ClinicName string `json:"clinic_name" binding:"required"`

	Diagnosis string `json:"diagnosis"`

	NextAppointment string `json:"next_appointment"`

	Veterinarian string `json:"veterinarian" binding:"required"`

	Treatment string `json:"treatment"`

	Type string `json:"type" binding:"required"`

	Medication string `json:"medication"`

	Description string `json:"description"`

	RecoveryStatus string `json:"recoveryStatus"`
}