package services

import (
	"errors"
	"gorm.io/gorm" 

	config "k9-system/config"

	"k9-system/constants"

	dto "k9-system/dto/k9_health"

	K9HealthRecordModel "k9-system/models/k9_health"

	"k9-system/utils"

	K9ProfileModel "k9-system/models/k9_profile"
)

func CreateHealthRecord(
	req dto.CreateK9HealthRecordRequest,
	) (*K9HealthRecordModel.HealthRecord, error) {

	tx := config.DB.Begin()

	visitedDate, err := utils.ParseDate(
		req.VisitedDate,
	)

	if err != nil {
		tx.Rollback()
		return nil, errors.New(
			constants.InvalidDate("VisitedDate"),
		)
	}

	releasedDate, err := utils.ParseDate(
		req.ReleasedDate,
	)

	if err != nil {
		tx.Rollback()
		return nil, errors.New(
			constants.InvalidDate("ReleasedDate"),
		)
	}

	nextAppointment, err := utils.ParseDate(
		req.NextAppointment,
	)

	if err != nil {
		tx.Rollback()
		return nil, errors.New(
			constants.InvalidDate("NextAppointment"),
		)
	}

	k9ProfileUUID, err := utils.ParseUUID(
		req.K9ID,
	)

	if err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.InvalidUUID("K9 Profile ID"),
		)
	}

	K9HealthRecord := K9HealthRecordModel.HealthRecord{

		K9ID: k9ProfileUUID,

		VisitedDate: visitedDate,

		ReleasedDate: releasedDate,

		ClinicName: req.ClinicName,

		Diagnosis: req.Diagnosis,

		NextAppointment: nextAppointment,

		Veterinarian: req.Veterinarian,

		Treatment: req.Treatment,

		Type: req.Type,

		Medication: req.Medication,

		Description: req.Description,

		RecoveryStatus: req.RecoveryStatus,
	}

	if err := tx.Create(&K9HealthRecord).Error; err != nil {

		tx.Rollback()

		return nil, err
	}

	if err := tx.
		Preload("K9Profile").
		First(
			&K9HealthRecord,
			"id = ?",
			K9HealthRecord.ID,
		).Error; err != nil {

		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return &K9HealthRecord, nil
}

func UpdateHealthRecord(
	id string,
	req dto.UpdateK9HealthRecordRequest,
) (*K9HealthRecordModel.HealthRecord, error) {

	var healthRecord K9HealthRecordModel.HealthRecord

	if err := config.DB.
		First(&healthRecord, "id = ?", id).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			return nil, errors.New(
				constants.NotFound("Health record"),
			)
		}

		return nil, errors.New(
			constants.FailedToRetrieve("health record"),
		)
	}

	tx := config.DB.Begin()

	if req.K9ID != nil {

		k9UUID, err := utils.ParseUUID(
			*req.K9ID,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidUUID("K9 Profile ID"),
			)
		}

		var k9Profile K9ProfileModel.K9Profile

		if err := tx.
			Where("id = ?", k9UUID).
			First(&k9Profile).
			Error; err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.NotFound("K9 profile"),
			)
		}

		healthRecord.K9ID = k9UUID
	}

	if req.VisitedDate != nil {

		visitedDate, err := utils.ParseDate(
			*req.VisitedDate,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidDate("VisitedDate"),
			)
		}

		healthRecord.VisitedDate = visitedDate
	}

	if req.ReleasedDate != nil {

		releasedDate, err := utils.ParseDate(
			*req.ReleasedDate,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidDate("ReleasedDate"),
			)
		}

		healthRecord.ReleasedDate = releasedDate
	}

	if req.NextAppointment != nil {

		nextAppointment, err := utils.ParseDate(
			*req.NextAppointment,
		)

		if err != nil {

			tx.Rollback()

			return nil, errors.New(
				constants.InvalidDate("NextAppointment"),
			)
		}

		healthRecord.NextAppointment = nextAppointment
	}

	if req.ClinicName != nil {
		healthRecord.ClinicName = *req.ClinicName
	}

	if req.Diagnosis != nil {
		healthRecord.Diagnosis = *req.Diagnosis
	}

	if req.Veterinarian != nil {
		healthRecord.Veterinarian = *req.Veterinarian
	}

	if req.Treatment != nil {
		healthRecord.Treatment = *req.Treatment
	}

	if req.Type != nil {
		healthRecord.Type = *req.Type
	}

	if req.Medication != nil {
		healthRecord.Medication = *req.Medication
	}

	if req.Description != nil {
		healthRecord.Description = *req.Description
	}

	if req.RecoveryStatus != nil {
		healthRecord.RecoveryStatus = *req.RecoveryStatus
	}

	if err := tx.Save(&healthRecord).Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToUpdate("health record"),
		)
	}

	if err := tx.
		Preload("K9Profile").
		First(
			&healthRecord,
			"id = ?",
			healthRecord.ID,
		).Error; err != nil {

		tx.Rollback()

		return nil, errors.New(
			constants.FailedToRetrieve("health record"),
		)
	}

	if err := tx.Commit().Error; err != nil {

		return nil, err
	}

	return &healthRecord, nil
}

func GetHealthRecords(
	search string,
	status string,
	healthRecordType string,
	pagination utils.Pagination,
) ([]K9HealthRecordModel.HealthRecord, utils.Pagination, error) {

	var total int64

	var healthRecords []K9HealthRecordModel.HealthRecord

	query := config.DB.Model(
		&K9HealthRecordModel.HealthRecord{},
	)

	query = query.Joins(`
		LEFT JOIN k9_profiles
		ON k9_profiles.id = health_records.k9_id
	`)

	if search != "" {

		searchPattern := "%" + search + "%"

		query = query.Where(`
			k9_profiles.name ILIKE ?
			OR health_records.veterinarian ILIKE ?
			OR health_records.diagnosis ILIKE ?
		`,
			searchPattern,
			searchPattern,
			searchPattern,
		)
	}

	if status != "" {
		query = query.Where(
			"health_records.status = ?",
			status,
		)
	}

	if healthRecordType != "" {
		query = query.Where(
			"health_records.type = ?",
			healthRecordType,
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, pagination, err
	}

	if err := query.
		Offset(pagination.Offset).
		Limit(pagination.Limit).
		Order("health_records.created_at DESC").
		Preload("K9Profile").
		Find(&healthRecords).
		Error; err != nil {

		return nil, pagination, err
	}

	pagination = utils.BuildPagination(
		pagination,
		total,
	)

	return healthRecords, pagination, nil
}

func GetaHealthRecordByID(
	id string,
) (*K9HealthRecordModel.HealthRecord, error) {

	var K9HealthRecord K9HealthRecordModel.HealthRecord

	if err := config.DB.Preload("K9Profile").First(&K9HealthRecord, "id = ?", id).Error; err != nil {
		return nil, err		

	}

	return &K9HealthRecord, nil
}