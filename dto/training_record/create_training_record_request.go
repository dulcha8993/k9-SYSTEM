package training_record

type CreateTrainingRecordRequest struct{

	K9ID string `json:"k9_id" binding:"required"`

	TrainerID string `json:"trainer_id" binding:"required"`

	TrainerName string `json:"trainer_name" binding:"required"`

	TrainingType string `json:"training_type" binding:"required"`

	TrainingModule string `json:"training_module" binding:"required"`

	StartDate string `json:"start_date" binding:"required"`

	EndDate string `json:"end_date" binding:"required"`

	TrainingCenterName string `json:"training_center_name" binding:"required"`

	TrainingCenterAddressLine1 string `json:"training_center_address_line_1" binding:"required"`
	
	TrainingCenterAddressLine2 string `json:"training_center_address_line_2" binding:"required"`

	TrainingCenterCountry string `json:"training_center_country" binding:"required"`

	PerformanceScore float64 `json:"performance_score"`

	Notes string `json:"notes"`

	Objectives string `json:"objectives"`
}
