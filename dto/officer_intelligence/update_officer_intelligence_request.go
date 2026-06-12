package officer_intelligence

type UpdateCriminalCaseRequest struct {
	
	CaseNumber *string `json:"case_number"`

	Title *string `json:"title"`

	Description *string `json:"description"`

	Outcome *string `json:"outcome"`
	
	Status *int `json:"status"`

	OpenDate *string `json:"open_date"`
	
	ClosedDate *string `json:"closed_date"`

	Role *string `json:"role"`

	Location *string `json:"location"`

	ReportingOfficer *string `json:"reporting_officer"`
}