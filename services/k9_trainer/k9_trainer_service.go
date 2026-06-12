package services

import (
	"errors"

	config "k9-system/config"

	"k9-system/constants"

	dto "k9-system/dto/k9_trainer"

	"k9-system/utils"

	K9Trainer "k9-system/models/k9_trainer"
)

func CreateK9Trainer(
	req dto.K9TrainerRequest,
	) (*dto.K9TrainerResponse, error){

	JoinDate, err := utils.ParseDate(
		req.JoinDate,
	)

	if err != nil {
		return nil, errors.New(
			constants.InvalidDate("JoinDate"),
		)
	}

	trainer := K9Trainer.K9Trainer{		

		Name: req.Name,

		BadgeNumber: req.BadgeNumber,

		Rank: req.Rank,

		Email: req.Email,

		Unit: req.Unit,

		JoinDate: JoinDate,

		ExperienceYears: req.ExperienceYears,

		Status: req.Status,
	}

		// Start transaction
		tx := config.DB.Begin()

		// Create trainer record

		if err := tx.Create(&trainer).Error; err != nil {

			tx.Rollback()

			return nil, err
		}

			// Create certifications if provided

		for _, cert := range req.Certifications {

	
			parsedDateObtained, err := utils.ParseDate(
				cert.DateObtained,
			)

			if err != nil {

				tx.Rollback()
				return nil, errors.New(
					constants.InvalidDate("DateObtained"),
				)
			}

		// Optional expiration date

			parsedExpDate, err := utils.ParseDate(
				cert.DateExpires,
			)

			if err != nil {

				tx.Rollback()

				return nil, errors.New(
					constants.InvalidDate("DateExpires"),
				)
			}

			certRecord := K9Trainer.K9TrainerCertification{

				TrainerID: trainer.ID,

				Name: cert.CertificationName,

				IssuingOrganization: cert.IssuingOrganization,

				DateObtained: parsedDateObtained,

				DateExpires: parsedExpDate,
			}

			if err := tx.Create(&certRecord).Error; err != nil {

				tx.Rollback()
				return nil, err
			}
		}

			tx.Commit()

			trainerResponse := dto.K9TrainerResponse{

				GenaralInfo: trainer,

				Certifications: []dto.K9TrainerCertificationResponse{},
			}

			for _, cert := range req.Certifications {

				certResponse := dto.K9TrainerCertificationResponse{
					Name: cert.CertificationName,
					IssuingOrganization: cert.IssuingOrganization,
					DateObtained: cert.DateObtained,
					DateExpires: cert.DateExpires,
				}
				trainerResponse.Certifications = append(trainerResponse.Certifications, certResponse)
			}

			return &trainerResponse, nil

}

func UpdateK9Trainer(
	id string,
	req dto.K9TrainerRequest,
) (*K9Trainer.K9Trainer, error) {

	var trainer K9Trainer.K9Trainer

		// Start transaction
	tx := config.DB.Begin()

	if err := tx.First(&trainer, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, errors.New(
			constants.NotFound("K9 Trainer"),
		)
	}

	if req.Name != "" {
		trainer.Name = req.Name
	}
	
	if req.BadgeNumber != "" {
		trainer.BadgeNumber = req.BadgeNumber
	}		

	if req.Rank != "" {
		trainer.Rank = req.Rank
	}

	if req.Email != "" {
		trainer.Email = req.Email
	}

	if req.Unit != "" {
		trainer.Unit = req.Unit
	}

	if req.JoinDate != "" {
		JoinDate, err := utils.ParseDate(
		req.JoinDate,
		)

		if err != nil {
			return nil, errors.New(
				constants.InvalidDate("JoinDate"),
			)
		}

		trainer.JoinDate = JoinDate
	}

	if req.ExperienceYears != 0 {
		trainer.ExperienceYears = req.ExperienceYears
	}

	if req.Status != 0 {
		trainer.Status = req.Status
	}

	if err := tx.Save(&trainer).Error; err != nil {
		tx.Rollback()

		return nil, errors.New(
			constants.FailedToUpdate("trainer"),
		)
	}

		// Replace additional certifications only if provided

	if req.Certifications != nil {

		// Delete existing certifications
		if err := tx.Where("trainer_id = ?", trainer.ID).Delete(&K9Trainer.K9TrainerCertification{}).Error; err != nil {
			tx.Rollback()
			return nil, errors.New(
				constants.FailedToUpdate("trainer - Cert"),
			)
		}

		// Add new certifications
		for _, cert := range req.Certifications {

			parsedDateObtained, err := utils.ParseDate(
				cert.DateObtained,
			)

			if err != nil {
				return nil, errors.New(
					constants.InvalidDate("DateObtained"),
				)
			}


			parsedExpDate, err := utils.ParseDate(
				cert.DateExpires,
			)

			if err != nil {
				return nil, errors.New(
					constants.InvalidDate("DateExpires"),
				)
			}


			certRecord := K9Trainer.K9TrainerCertification{

				TrainerID: trainer.ID,

				Name: cert.CertificationName,

				IssuingOrganization: cert.IssuingOrganization,

				DateObtained: parsedDateObtained,

				DateExpires: parsedExpDate,
			}

			if err := tx.Create(&certRecord).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &trainer, nil

}

func GetK9Trainers(
	search string,
	status string,
	trainer_id string,
	pagination utils.Pagination,
) ([]K9Trainer.K9Trainer, utils.Pagination, error) {

	var trainers []K9Trainer.K9Trainer

	var total int64

	query := config.DB.Model(&K9Trainer.K9Trainer{})

	if search != "" {

		searchPattern := "%" + search + "%"

		query = query.Where(`name ILIKE ? OR badge_number ILIKE ? OR rank ILIKE ? OR email ILIKE ? OR unit ILIKE ?`,
		 searchPattern, searchPattern, searchPattern, searchPattern, searchPattern)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if trainer_id != "" {
		query = query.Where("id = ?", trainer_id)
	}

	query.Count(&total)

	query = query.Offset(pagination.Offset).Limit(pagination.Limit).Order("created_at DESC")

	pagination = utils.BuildPagination(
		pagination,
		total,
	)

	if err := query.Preload("Certifications").Find(&trainers).Error; err != nil {
		return nil, pagination, err
	}

		return trainers, pagination, nil

}

func AssignK9TrainerToK9(
	req dto.CreateK9TrainerAssignRequest,
) (*K9Trainer.K9TrainerHistory, error) {

	trainerUUID, err := utils.ParseUUID(
		req.TrainerID,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID("Trainer ID"),
		)
	}

	k9ProfileUUID, err := utils.ParseUUID(
		req.K9ID,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID("K9 Profile ID"),
		)
	}

	startDate, err := utils.ParseDate(
		req.StartDate,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidDate("StartDate"),
		)
	}

	endDate, err := utils.ParseDate(
		req.EndDate,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidDate("EndDate"),
		)
	}

	tx := config.DB.Begin()

	trainerHistory := K9Trainer.K9TrainerHistory{

		TrainerID: trainerUUID,

		K9ID: &k9ProfileUUID,

		StartDate: startDate,

		EndDate: endDate,

		Status: 1,
	}

	// Deactivate previous assignments

	if err := tx.
		Model(&K9Trainer.K9TrainerHistory{}).
		Where(
			"trainer_id = ?",
			trainerUUID,
		).
		Update(
			"status",
			0,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToUpdate(
				"previous trainer history",
			),
		)
	}

	// Create new assignment

	if err := tx.
		Create(&trainerHistory).
		Error; err != nil {

		tx.Rollback()

		return nil, err
	}

	if err := tx.
		Preload("Trainer").
		Preload("K9Profile").
		First(
			&trainerHistory,
			"id = ?",
			trainerHistory.ID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToRetrieve(
				"trainer assignment",
			),
		)
	}

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &trainerHistory, nil
}