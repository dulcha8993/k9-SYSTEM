package utils

import (
	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func IsValidUUID(id string) bool {

	_, err := uuid.Parse(id)

	return err == nil
}