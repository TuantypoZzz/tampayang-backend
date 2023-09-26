package sanitizeLib

import (
	"fmt"
	"regexp"
	"strings"
)

func PhoneNumber(phoneNumber string) string {
	// Define a regular expression pattern to match non-numeric characters
	regex := regexp.MustCompile("[^0-9]+")

	// Use the ReplaceAllString function to remove non-numeric characters
	cleanPhone := regex.ReplaceAllString(phoneNumber, "")

	if cleanPhone != "" {
		// Check if the cleaned phone number starts with "0"
		if strings.HasPrefix(cleanPhone, "0") {
			// Add the country code "+62" and remove the leading "0"
			cleanPhone = "+62" + cleanPhone[1:]
		} else {
			// Add a generic country code prefix "+"
			cleanPhone = "+" + cleanPhone
		}
	}

	// Print the cleaned phone number
	fmt.Println("Cleaned Phone:", cleanPhone)

	// Return cleaned phone number
	return ""

}