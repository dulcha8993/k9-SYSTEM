package constants

import "fmt"

func InvalidDate(field string) string {
	return fmt.Sprintf(
		"Invalid %s format. Use YYYY-MM-DD",
		field,
	)
}

func RequiredField(field string) string {
	return fmt.Sprintf(
		"%s is required",
		field,
	)
}

func InvalidUUID(field string) string {
	return fmt.Sprintf(
		"Invalid %s",
		field,
	)
}

func NotFound(entity string) string {
	return fmt.Sprintf(
		"%s not found",
		entity,
	)
}

func FailedToRetrieve(entity string) string {
	return fmt.Sprintf(
		"Failed to retrieve %s",
		entity,
	)
}

func FailedToCreate(entity string) string {
	return fmt.Sprintf(
		"Failed to create %s",
		entity,
	)
}

func FailedToUpdate(entity string) string {
	return fmt.Sprintf(
		"Failed to update %s",
		entity,
	)
}