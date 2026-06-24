package dashboard

import (
	K9ProfileModel "k9-system/models/k9_profile"

	K9HealthRecordModel "k9-system/models/k9_health"

	"github.com/google/uuid"
)

type DashboardResponse struct {
	ActiveK9Profiles        int64                      `json:"active_k9_profiles"`
	ActiveHandlers          int64                      `json:"active_handlers"`
	NewK9ProfilesThisMonth  int64                      `json:"new_k9_profiles_this_month"`
	AlertsCount				int64                      `json:"alerts_count"`

	Alerts                  []string                   `json:"alerts"`

	TrainingSessionsByK9    []TrainingSessionSummary   `json:"training_sessions_by_k9"`

	K9Activities            []K9ActivitySummary        `json:"k9_activities"`

	RecentK9Profiles        []K9ProfileModel.K9Profile `json:"recent_k9_profiles"`

	UpcomingAppointments    []K9HealthRecordModel.HealthRecord `json:"upcoming_appointments"`
}

type TrainingSessionSummary struct {
	K9ID          uuid.UUID `json:"k9_id"`
	K9Name        string    `json:"k9_name"`
	TrainingCount int64     `json:"training_count"`
}

type K9ActivitySummary struct {
	K9ID                uuid.UUID `json:"k9_id"`
	K9Name              string    `json:"k9_name"`
	TrainingRecords     int64     `json:"training_records"`
	HealthRecords       int64     `json:"health_records"`
	CriminalCases       int64     `json:"criminal_cases"`
	AlertsCount         int64     `json:"alerts_count"`
	TotalActivities      int64     `json:"total_activities"`
}