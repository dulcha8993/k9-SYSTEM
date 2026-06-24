package lookup

type LookupResponse struct {
	Departments			   []string `json:"departments"`

	Breeds                 []string `json:"breeds"`
	Colors                 []string `json:"colors"`

	Ranks                  []string `json:"ranks"`
	Units                  []string `json:"units"`

	ClinicNames            []string `json:"clinic_names"`
	Veterinarians          []string `json:"veterinarians"`
	Diagnoses              []string `json:"diagnoses"`
	HealthRecordTypes      []string `json:"health_record_types"`

	TrainingTypes          []string `json:"training_types"`
	TrainingModules        []string `json:"training_modules"`
	TrainingCenters        []string `json:"training_centers"`
	TrainingCountries      []string `json:"training_countries"`

	CriminalCaseTitles     []string `json:"criminal_case_titles"`
	CriminalCaseTypes      []string `json:"criminal_case_types"`
	CriminalCaseLocations  []string `json:"criminal_case_locations"`
}