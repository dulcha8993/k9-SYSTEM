package services

import (
	"errors"
	// "gorm.io/gorm" 

	config "k9-system/config"

	"k9-system/constants"

	dto "k9-system/dto/training_record"

	TrainingRecord "k9-system/models/training_record"

	"k9-system/utils"

	K9TrainerModel "k9-system/models/k9_trainer"

	K9ProfileModel "k9-system/models/k9_profile"
)

func CreateTrainingRecord(
	req dto.CreateTrainingRecordRequest,
) (*TrainingRecord.TrainingRecord, error) {

	startDate, err := utils.ParseDate(
		req.StartDate,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidDate(
				"StartDate",
			),
		)
	}

	endDate, err := utils.ParseDate(
		req.EndDate,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidDate(
				"EndDate",
			),
		)
	}

	trainerUUID, err := utils.ParseUUID(
		req.TrainerID,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID(
				"Trainer ID",
			),
		)
	}

	k9ProfileUUID, err := utils.ParseUUID(
		req.K9ID,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID(
				"K9 Profile ID",
			),
		)
	}

	tx := config.DB.Begin()

	// Optional validation that trainer exists

	var trainer K9TrainerModel.K9Trainer

	if err := tx.
		First(
			&trainer,
			"id = ?",
			trainerUUID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.NotFound(
				"trainer",
			),
		)
	}

	// Optional validation that K9 exists

	var k9Profile K9ProfileModel.K9Profile

	if err := tx.
		First(
			&k9Profile,
			"id = ?",
			k9ProfileUUID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.NotFound(
				"K9 profile",
			),
		)
	}

	trainingRecord := TrainingRecord.TrainingRecord{

		K9ID: k9ProfileUUID,

		TrainerID: trainerUUID,

		TrainingType: req.TrainingType,

		TrainingModule: req.TrainingModule,

		StartDate: *startDate,

		EndDate: *endDate,

		TrainingCenterName: req.TrainingCenterName,

		TrainingCenterAddressLine1: req.TrainingCenterAddressLine1,

		TrainingCenterAddressLine2: req.TrainingCenterAddressLine2,

		TrainingCenterCountry: req.TrainingCenterCountry,

		PerformanceScore: req.PerformanceScore,

		Notes: req.Notes,

		Objectives: req.Objectives,
	}

	if err := tx.
		Create(&trainingRecord).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToCreate(
				"training record",
			),
		)
	}

	if err := tx.
		Preload("K9Profile").
		Preload("K9Trainer").
		First(
			&trainingRecord,
			"id = ?",
			trainingRecord.ID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToRetrieve(
				"training record",
			),
		)
	}

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &trainingRecord, nil
}

func UpdateTrainingRecord(
	id string,
	req dto.UpdateTrainingRecordRequest,
) (*TrainingRecord.TrainingRecord, error) {

	tx := config.DB.Begin()

	var trainingRecord TrainingRecord.TrainingRecord

	if err := tx.
		First(
			&trainingRecord,
			"id = ?",
			id,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.NotFound(
				"Training Record",
			),
		)
	}

	if req.K9ID != nil {

		k9UUID, err := utils.ParseUUID(
			*req.K9ID,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidUUID(
					"K9 ID",
				),
			)
		}

		trainingRecord.K9ID = k9UUID
	}

	if req.TrainerID != nil {

		trainerUUID, err := utils.ParseUUID(
			*req.TrainerID,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidUUID(
					"Trainer ID",
				),
			)
		}

		trainingRecord.TrainerID = trainerUUID
	}

	if req.TrainerName != nil {
		trainingRecord.TrainerName = *req.TrainerName
	}

	if req.TrainingType != nil {
		trainingRecord.TrainingType = *req.TrainingType
	}

	if req.TrainingModule != nil {
		trainingRecord.TrainingModule = *req.TrainingModule
	}

	if req.StartDate != nil {

		startDate, err := utils.ParseDate(
			*req.StartDate,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidDate(
					"StartDate",
				),
			)
		}

		trainingRecord.StartDate = *startDate
	}

	if req.EndDate != nil {

		endDate, err := utils.ParseDate(
			*req.EndDate,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidDate(
					"EndDate",
				),
			)
		}

		trainingRecord.EndDate = *endDate
	}

	if req.TrainingCenterName != nil {
		trainingRecord.TrainingCenterName = *req.TrainingCenterName
	}

	if req.TrainingCenterAddressLine1 != nil {
		trainingRecord.TrainingCenterAddressLine1 = *req.TrainingCenterAddressLine1
	}

	if req.TrainingCenterAddressLine2 != nil {
		trainingRecord.TrainingCenterAddressLine2 = *req.TrainingCenterAddressLine2
	}

	if req.TrainingCenterCountry != nil {
		trainingRecord.TrainingCenterCountry = *req.TrainingCenterCountry
	}

	if req.PerformanceScore != nil {
		trainingRecord.PerformanceScore = *req.PerformanceScore
	}

	if req.Notes != nil {
		trainingRecord.Notes = *req.Notes
	}

	if req.Objectives != nil {
		trainingRecord.Objectives = *req.Objectives
	}

	if err := tx.
		Save(&trainingRecord).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToUpdate(
				"training record",
			),
		)
	}

	if err := tx.
		Preload("K9Profile").
		Preload("K9Trainer").
		First(
			&trainingRecord,
			"id = ?",
			trainingRecord.ID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToRetrieve(
				"training record",
			),
		)
	}

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &trainingRecord, nil
}

func GetTrainingRecords(
	search string,
	status string,
	trainingRecordID string,
	pagination utils.Pagination,
) ([]TrainingRecord.TrainingRecord, utils.Pagination, error) {

	var trainingRecords []TrainingRecord.TrainingRecord

	var total int64

	query := config.DB.Model(
		&TrainingRecord.TrainingRecord{},
	)

	if search != "" {

		searchPattern := "%" + search + "%"

		query = query.Where(`
			training_type ILIKE ?
			OR training_module ILIKE ?
			OR trainer_name ILIKE ?
		`,
			searchPattern,
			searchPattern,
			searchPattern,
		)
	}

	if status != "" {

		query = query.Where(
			"status = ?",
			status,
		)
	}

	if trainingRecordID != "" {

		query = query.Where(
			"id = ?",
			trainingRecordID,
		)
	}

	if err := query.Count(&total).Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve(
				"training records",
			),
		)
	}

	if err := query.
		Offset(pagination.Offset).
		Limit(pagination.Limit).
		Order("created_at DESC").
		Preload("K9Profile").
		Preload("K9Trainer").
		Find(&trainingRecords).
		Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve(
				"training records",
			),
		)
	}

	pagination = utils.BuildPagination(
		pagination,
		total,
	)

	return trainingRecords, pagination, nil
}