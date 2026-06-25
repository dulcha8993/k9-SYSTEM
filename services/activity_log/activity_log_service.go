package activity_log

import (
	"encoding/json"

	config "k9-system/config"

	activityLogDTO "k9-system/dto/activity_log"
	ActivityLogModel "k9-system/models/activity_log"

	"github.com/google/uuid"

	"k9-system/constants"
)

func LogActivity(
	userID uuid.UUID,
	module string,
	action string,
	recordID *uuid.UUID,
	status string,
	description string,
	errorMessage *string,
	changes []activityLogDTO.FieldChange,
) error {

	var changesJSON []byte

	if changes != nil {

		data, err := json.Marshal(changes)
		if err != nil {
			return err
		}

		changesJSON = data
	}

	log := ActivityLogModel.ActivityLog{

		UserID: userID,

		Module: module,

		Action: action,

		RecordID: recordID,

		Status: status,

		Description: description,

		ErrorMessage: errorMessage,

		Changes: changesJSON,
	}

	return config.DB.Create(&log).Error
}

func LogSuccess(
	userID uuid.UUID,
	module string,
	action string,
	recordID *uuid.UUID,
	description string,
	changes []activityLogDTO.FieldChange,
) error {

	return LogActivity(
		userID,
		module,
		action,
		recordID,
		constants.ActivitySuccess,
		description,
		nil,
		changes,
	)
}

func LogFailure(
	userID uuid.UUID,
	module string,
	action string,
	description string,
	err error,
) error {

	var errorMessage *string

	if err != nil {

		msg := err.Error()

		errorMessage = &msg
	}

	return LogActivity(
		userID,
		module,
		action,
		nil,
		constants.ActivityFailed,
		description,
		errorMessage,
		nil,
	)
}