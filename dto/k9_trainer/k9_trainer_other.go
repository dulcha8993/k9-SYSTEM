package k9_trainer

type K9TrainerCertificationRequest struct {

	CertificationName string `json:"certification_name"`

	IssuingOrganization string `json:"issuing_organization"`

	DateObtained string `json:"date_obtained"`

	DateExpires string `json:"date_expires"`
}

type K9TrainerCertificationResponse struct {

	Name string `json:"name"`

	IssuingOrganization string `json:"issuing_organization"`

	DateObtained string `json:"date_obtained"`

	DateExpires string `json:"date_expires"`
}

type CreateK9TrainerAssignRequest struct {

	TrainerID string `json:"trainer_id"`

	K9ID string `json:"k9_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`
}