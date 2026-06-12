package utils

func SafeString(value *string) string {

	if value == nil {
		return ""
	}

	return *value
}