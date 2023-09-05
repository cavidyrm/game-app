package phonenumber

import "strconv"

func IsValid(phoneNumber string) bool {
	// TODO - use regular expression to validate phone number to check +98 prefix
	if len(phoneNumber) != 11 {
		return false
	}
	if phoneNumber[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phoneNumber[2:]); err == nil {
		return false
	}
	return true
}
