package officer_intelligence

import (
	"time"

	"github.com/google/uuid"
)

type CriminalCaseResponse struct {

	ID uuid.UUID `json:"id"`

	K9TrainerHistoryID uuid.UUID `json:"trainer_history_id"`

	TrainerName string `json:"trainer_name"`

	TrainerID uuid.UUID `json:"trainer_id"`

	TrainerBadgeNumber *string `json:"trainer_badge_number"`

	K9ID *uuid.UUID `json:"k9_id"`

	K9Name *string `json:"k9_name"`

	K9Code *string `json:"k9_code"`

	MicrochipNumber *string `json:"microchip_number"`

	CaseNumber string `json:"case_number"`

	Title string `json:"title"`

	Description string `json:"description"`		

	Outcome string `json:"outcome"`
	
	Status int `json:"status"`

	OpenDate time.Time `json:"open_date"`
	
	ClosedDate *time.Time `json:"closed_date"`

	Role string `json:"role"`
	
	Location string `json:"location"`

	ReportingOfficer string `json:"reporting_officer"`

	CreatedAt time.Time `json:"created_at"`
	
	UpdatedAt time.Time `json:"updated_at"`
}
