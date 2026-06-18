package services

import (
	"time"

	config "k9-system/config"

	dto "k9-system/dto/dashboard"

	K9HealthRecordModel "k9-system/models/k9_health"

	K9ProfileModel "k9-system/models/k9_profile"

	K9TrainerModel "k9-system/models/k9_trainer"

	TrainingRecordModel "k9-system/models/training_record"
)

func GetDashboard() (*dto.DashboardResponse, error) {

	var result dto.DashboardResponse

	// ===============================
	// Active K9 Profiles
	// ===============================

	config.DB.
		Model(&K9ProfileModel.K9Profile{}).
		Where("status = ?", 1).
		Count(&result.ActiveK9Profiles)

	// ===============================
	// Active Handlers
	// ===============================

	config.DB.
		Model(&K9TrainerModel.K9Trainer{}).
		Where("status = ?", 1).
		Count(&result.ActiveHandlers)

	// ===============================
	// New K9 Profiles This Month
	// ===============================

	startOfMonth := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		1,
		0,
		0,
		0,
		0,
		time.Now().Location(),
	)

	config.DB.
		Model(&K9ProfileModel.K9Profile{}).
		Where("created_at >= ?", startOfMonth).
		Count(&result.NewK9ProfilesThisMonth)

	// ===============================
	// Alerts
	// ===============================

	result.Alerts = []string{}

	// ===============================
	// Recent K9 Profiles
	// ===============================

	config.DB.
		Order("created_at DESC").
		Limit(3).
		Find(&result.RecentK9Profiles)

	// ===============================
	// Upcoming Appointments
	// ===============================

	config.DB.
		Preload("K9Profile").
		Where("next_appointment IS NOT NULL").
		Where("next_appointment >= ?", time.Now()).
		Order("next_appointment ASC").
		Limit(4).
		Find(&result.UpcomingAppointments)

	// ===============================
	// Training Sessions By K9
	// ===============================

	var trainingSummary []dto.TrainingSessionSummary

	config.DB.
		Table("training_records").
		Select(`
			k9_profiles.id as k9_id,
			k9_profiles.name as k9_name,
			count(training_records.id) as training_count
		`).
		Joins(`
			JOIN k9_profiles
			ON k9_profiles.id = training_records.k9_id
		`).
		Group("k9_profiles.id, k9_profiles.name").
		Scan(&trainingSummary)

	result.TrainingSessionsByK9 = trainingSummary

	// ===============================
	// K9 Activities
	// ===============================

	var activities []dto.K9ActivitySummary

	var k9Profiles []K9ProfileModel.K9Profile

	config.DB.Find(&k9Profiles)

	for _, k9 := range k9Profiles {

		var trainingCount int64
		var healthCount int64
		var caseCount int64

		config.DB.
			Model(&TrainingRecordModel.TrainingRecord{}).
			Where("k9_id = ?", k9.ID).
			Count(&trainingCount)

		config.DB.
			Model(&K9HealthRecordModel.HealthRecord{}).
			Where("k9_id = ?", k9.ID).
			Count(&healthCount)

		config.DB.
			Table("criminal_cases").
			Joins(`
				JOIN k9_trainer_histories
				ON criminal_cases.k9_trainer_history_id = k9_trainer_histories.id
			`).
			Where("k9_trainer_histories.k9_id = ?", k9.ID).
			Count(&caseCount)

		activities = append(
			activities,
			dto.K9ActivitySummary{
				K9ID:            k9.ID,
				K9Name:          k9.Name,
				TrainingRecords: trainingCount,
				HealthRecords:   healthCount,
				CriminalCases:   caseCount,
				AlertsCount: 0,
				TotalActivities: trainingCount + healthCount + caseCount,
			},
		)
	}

	result.K9Activities = activities

	return &result, nil
}