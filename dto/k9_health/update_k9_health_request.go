package k9_health

type UpdateK9HealthRecordRequest struct {
	K9ID *string `json:"k9_id"`

	VisitedDate *string `json:"visited_date"` 

	ReleasedDate *string `json:"released_date"`

	ClinicName *string `json:"clinic_name"`

	Diagnosis *string `json:"diagnosis"`

	NextAppointment *string `json:"next_appointment"`

	Veterinarian *string `json:"veterinarian"`

	Treatment *string `json:"treatment"`

	Type *string `json:"type"`

	Medication *string `json:"medication"`

	Description *string `json:"description"`

	RecoveryStatus *string `json:"recoveryStatus"`
}