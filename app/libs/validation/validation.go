package validationLib

import (
	sanitizeLib "tampayang-backend/app/libs/sanitize"
	globalFunction "tampayang-backend/core/functions"
)

func IsValidPhoneNumber(phoneNumber string) bool {

	phoneNumber = sanitizeLib.PhoneNumber(phoneNumber)

	if globalFunction.IsEmpty(phoneNumber) {
		return false
	}

	if len(phoneNumber) < 11 || len(phoneNumber) > 15 {
		return false
	}

	return true
}
