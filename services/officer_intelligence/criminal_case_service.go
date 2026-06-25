package services 

import (
	"errors"
	"gorm.io/gorm" 
	"time"

	config "k9-system/config"

	"k9-system/constants"

	dto "k9-system/dto/officer_intelligence"

	CriminalCaseModel "k9-system/models/officer_intelligence"

	K9TrainerHistoryModel "k9-system/models/k9_trainer"

	"k9-system/utils"

	activityLogDTO "k9-system/dto/activity_log"
)

func CreateCriminalCase(
	req dto.CreateCriminalCaseRequest,
) (*CriminalCaseModel.CriminalCase, error) {

	trainerHistoryUUID, err := utils.ParseUUID(
		req.K9TrainerHistoryID,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID(
				"K9 Trainer History ID",
			),
		)
	}

	openDate, err := utils.ParseDate(
		req.OpenDate,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidDate(
				"OpenDate",
			),
		)
	}

	var closedDate *time.Time

	if req.ClosedDate != nil {

		cd, err := utils.ParseDate(
			*req.ClosedDate,
		)

		if err != nil {

			return nil, errors.New(
				constants.InvalidDate(
					"ClosedDate",
				),
			)
		}

		closedDate = cd
	}

	tx := config.DB.Begin()

	// Validate trainer history exists

	var trainerHistory K9TrainerHistoryModel.K9TrainerHistory

	if err := tx.
		First(
			&trainerHistory,
			"id = ?",
			trainerHistoryUUID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.NotFound(
				"K9 trainer history",
			),
		)
	}

	criminalCase := CriminalCaseModel.CriminalCase{

		K9TrainerHistoryID: trainerHistoryUUID,

		CaseNumber: req.CaseNumber,

		Title: req.Title,

		Description: req.Description,

		Outcome: req.Outcome,

		Status: req.Status,

		OpenDate: *openDate,

		ClosedDate: closedDate,

		Role: req.Role,

		Location: req.Location,

		ReportingOfficer: req.ReportingOfficer,
	}

	if err := tx.
		Create(&criminalCase).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToCreate(
				"criminal case",
			),
		)
	}

	if err := tx.
		Preload("TrainerHistory").
		Preload("TrainerHistory.Trainer").
		Preload("TrainerHistory.K9Profile").
		First(
			&criminalCase,
			"id = ?",
			criminalCase.ID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToRetrieve(
				"criminal case",
			),
		)
	}

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &criminalCase, nil
}

func UpdateCriminalCase(
	id string,
	req dto.UpdateCriminalCaseRequest,
) (*CriminalCaseModel.CriminalCase, []activityLogDTO.FieldChange, error) {

	criminalCaseUUID, err := utils.ParseUUID(
		id,
	)

	if err != nil {

		return nil, nil, errors.New(
			constants.InvalidUUID(
				"Criminal Case ID",
			),
		)
	}

	tx := config.DB.Begin()

	var criminalCase CriminalCaseModel.CriminalCase

	if err := tx.
		First(
			&criminalCase,
			"id = ?",
			criminalCaseUUID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, nil, errors.New(
			constants.NotFound(
				"criminal case",
			),
		)
	}

		// Keep original copy
	original := criminalCase

	if req.CaseNumber != nil {
		criminalCase.CaseNumber = *req.CaseNumber
	}

	if req.Title != nil {
		criminalCase.Title = *req.Title
	}

	if req.Description != nil {
		criminalCase.Description = *req.Description
	}

	if req.Outcome != nil {
		criminalCase.Outcome = *req.Outcome
	}

	if req.Status != nil {
		criminalCase.Status = *req.Status
	}

	if req.OpenDate != nil {

		openDate, err := utils.ParseDate(
			*req.OpenDate,
		)

		if err != nil {

			tx.Rollback()

			return nil, nil, errors.New(
				constants.InvalidDate(
					"OpenDate",
				),
			)
		}

		criminalCase.OpenDate = *openDate
	}

	if req.ClosedDate != nil {

		closedDate, err := utils.ParseDate(
			*req.ClosedDate,
		)

		if err != nil {

			tx.Rollback()

			return nil, nil, errors.New(
				constants.InvalidDate(
					"ClosedDate",
				),
			)
		}

		criminalCase.ClosedDate = closedDate
	}

	if req.Role != nil {
		criminalCase.Role = *req.Role
	}

	if req.Location != nil {
		criminalCase.Location = *req.Location
	}

	if req.ReportingOfficer != nil {
		criminalCase.ReportingOfficer = *req.ReportingOfficer
	}

	if err := tx.
		Save(&criminalCase).
		Error; err != nil {

		tx.Rollback()

		return nil, nil, errors.New(
			constants.FailedToUpdate(
				"criminal case",
			),
		)
	}

	if err := tx.
		Preload("TrainerHistory").
		Preload("TrainerHistory.Trainer").
		Preload("TrainerHistory.K9Profile").
		First(
			&criminalCase,
			"id = ?",
			criminalCase.ID,
		).
		Error; err != nil {

		tx.Rollback()

		return nil, nil, errors.New(
			constants.FailedToRetrieve(
				"criminal case",
			),
		)
	}

	// Detect changes BEFORE Save
	changes := utils.GetChanges(
		original,
		criminalCase,
	)

	if err := tx.Commit().Error; err != nil {

		return nil, nil, err
	}

	return &criminalCase, changes, nil
}

func GetCriminalCases(
	search string,
	caseType string,
	pagination utils.Pagination,
) ([]CriminalCaseModel.CriminalCase, utils.Pagination, error) {

	var total int64

	var criminalCases []CriminalCaseModel.CriminalCase

	query := config.DB.Model(
		&CriminalCaseModel.CriminalCase{},
	)

	if search != "" {

		likeSearch := "%" + search + "%"

		query = query.Where(
			config.DB.
				Where("case_number ILIKE ?", likeSearch).
				Or("title ILIKE ?", likeSearch).
				Or("outcome ILIKE ?", likeSearch).
				Or("location ILIKE ?", likeSearch).
				Or("reporting_officer ILIKE ?", likeSearch),
		)
	}

	if caseType != "" {

		query = query.Where(
			"case_type = ?",
			caseType,
		)
	}

	if err := query.Count(&total).Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve(
				"criminal cases",
			),
		)
	}

	if err := query.
		Offset(pagination.Offset).
		Limit(pagination.Limit).
		Order("created_at DESC").
		Preload("TrainerHistory").
		Preload("TrainerHistory.Trainer").
		Preload("TrainerHistory.K9Profile").
		Find(&criminalCases).
		Error; err != nil {

		return nil, pagination, errors.New(
			constants.FailedToRetrieve(
				"criminal cases",
			),
		)
	}

	pagination = utils.BuildPagination(
		pagination,
		total,
	)

	return criminalCases, pagination, nil
}

func GetCriminalCaseByID(
	id string,
) (*CriminalCaseModel.CriminalCase, error) {

	criminalCaseUUID, err := utils.ParseUUID(
		id,
	)

	if err != nil {

		return nil, errors.New(
			constants.InvalidUUID(
				"Criminal Case ID",
			),
		)
	}

	var criminalCase CriminalCaseModel.CriminalCase

	if err := config.DB.
		Preload("TrainerHistory").
		Preload("TrainerHistory.Trainer").
		Preload("TrainerHistory.K9Profile").
		First(
			&criminalCase,
			"id = ?",
			criminalCaseUUID,
		).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return nil, errors.New(
				constants.NotFound(
					"criminal case",
				),
			)
		}

		return nil, errors.New(
			constants.FailedToRetrieve(
				"criminal case",
			),
		)
	}

	return &criminalCase, nil
}