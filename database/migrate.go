package database

import (
	"fmt"
	"k9-system/config"
)

import (
	userModel "k9-system/models/user"
	accessModel "k9-system/models/access"
	K9ProfileModel "k9-system/models/k9_profile"
	K9TrainerModel "k9-system/models/k9_trainer"
	TrainingRecordModel "k9-system/models/training_record"
	K9HealthRecordModel "k9-system/models/k9_health"
	OfficerInterligenceModel "k9-system/models/officer_intelligence"
	ActivityLogModel "k9-system/models/activity_log"
)

func Migrate() {

	fmt.Println("Running migrations...")

	err := config.DB.AutoMigrate(
		&userModel.UserRole{},
		&userModel.User{},
		&accessModel.SystemModule{},
		&accessModel.ModuleAccess{},
		&accessModel.UserAdditionalModuleAccess{},
		&K9ProfileModel.K9Profile{},
		&K9TrainerModel.K9Trainer{},
		&K9TrainerModel.K9TrainerCertification{},
		&TrainingRecordModel.TrainingRecord{},
		&K9HealthRecordModel.HealthRecord{},
		&K9TrainerModel.K9TrainerHistory{},
		&OfficerInterligenceModel.CriminalCase{},
		&ActivityLogModel.ActivityLog{},

	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration completed")
}