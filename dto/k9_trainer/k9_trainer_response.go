package k9_trainer

import (
	K9Trainer "k9-system/models/k9_trainer"
)

type K9TrainerResponse struct {

	GenaralInfo K9Trainer.K9Trainer `json:"general_info"`

	Certifications []K9TrainerCertificationResponse `json:"certifications"`
}