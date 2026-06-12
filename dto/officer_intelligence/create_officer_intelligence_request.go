package officer_intelligence

type CreateCriminalCaseRequest struct {

	K9TrainerHistoryID string `json:"trainer_history_id" binding:"required"`

	CaseNumber string `json:"case_number" binding:"required"`

	Title string `json:"title" binding:"required"`

	Description string `json:"description"`
	
	Outcome string `json:"outcome"`

	Status int `json:"status" binding:"required"`
	
	OpenDate string `json:"open_date" binding:"required"`

	ClosedDate *string `json:"closed_date"`

	Role string `json:"role" binding:"required"`

	Location string `json:"location" binding:"required"`

	ReportingOfficer string `json:"reporting_officer" binding:"required"`

}