package validationLib

import (
	sanitizeLib "github.com/nulla-vis/golang-fiber-template/app/libs/sanitize"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
)

func IsValidPhoneNumber(phoneNumber string) bool {
	
	phoneNumber = sanitizeLib.PhoneNumber(phoneNumber)

	if globalFunction.IsEmpty(phoneNumber) {
		return false
	}

	if (len(phoneNumber) < 11 || len(phoneNumber) > 15) {
		return false
	}

	return true
}