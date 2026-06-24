package lookup

import (
	config "k9-system/config"

	dto "k9-system/dto/lookup"
)

func GetLookups() (
	*dto.LookupResponse,
	error,
) {

	result := dto.LookupResponse{}

	// Departments

	config.DB.
		Table("users").
		Distinct().
		Pluck(
			"department",
			&result.Departments,
		)

	// K9 Profiles

	config.DB.
		Table("k9_profiles").
		Distinct().
		Pluck(
			"breed",
			&result.Breeds,
		)

	config.DB.
		Table("k9_profiles").
		Distinct().
		Pluck(
			"color",
			&result.Colors,
		)

	// Trainers

	config.DB.
		Table("k9_trainers").
		Distinct().
		Pluck(
			"rank",
			&result.Ranks,
		)

	config.DB.
		Table("k9_trainers").
		Distinct().
		Pluck(
			"unit",
			&result.Units,
		)

	// Health Records

	config.DB.
		Table("health_records").
		Distinct().
		Pluck(
			"clinic_name",
			&result.ClinicNames,
		)

	config.DB.
		Table("health_records").
		Distinct().
		Pluck(
			"veterinarian",
			&result.Veterinarians,
		)

	config.DB.
		Table("health_records").
		Distinct().
		Pluck(
			"diagnosis",
			&result.Diagnoses,
		)

	config.DB.
		Table("health_records").
		Distinct().
		Pluck(
			"type",
			&result.HealthRecordTypes,
		)

	// Training Records

	config.DB.
		Table("training_records").
		Distinct().
		Pluck(
			"training_type",
			&result.TrainingTypes,
		)

	config.DB.
		Table("training_records").
		Distinct().
		Pluck(
			"training_module",
			&result.TrainingModules,
		)

	config.DB.
		Table("training_records").
		Distinct().
		Pluck(
			"training_center_name",
			&result.TrainingCenters,
		)

	config.DB.
		Table("training_records").
		Distinct().
		Pluck(
			"training_center_country",
			&result.TrainingCountries,
		)

	config.DB.
		Table("criminal_cases").
		Distinct().
		Pluck(
			"title",
			&result.CriminalCaseTitles,
		)

	config.DB.
		Table("criminal_cases").
		Distinct().
		Pluck(
			"role",
			&result.CriminalCaseTypes,
		)

	config.DB.
		Table("criminal_cases").
		Distinct().
		Pluck(
			"location",
			&result.CriminalCaseLocations,
		)

	return &result, nil
}