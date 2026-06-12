package training_record

type UpdateTrainingRecordRequest struct{
	K9ID *string `json:"k9_id"`

	TrainerID *string `json:"trainer_id"`

	TrainerName *string `json:"trainer_name"`

	TrainingType *string `json:"training_type"`

	TrainingModule *string `json:"training_module"`

	StartDate *string `json:"start_date"`

	EndDate *string `json:"end_date"`

	TrainingCenterName *string `json:"training_center_name"`

	TrainingCenterAddressLine1 *string `json:"training_center_address_line_1"`
	
	TrainingCenterAddressLine2 *string `jsonß:"training_center_address_line_2"`

	TrainingCenterCountry *string `json:"training_center_country"`

	PerformanceScore *float64 `json:"performance_score"`

	Notes *string `json:"notes"`

	Objectives *string `json:"objectives"`
}