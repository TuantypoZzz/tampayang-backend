package sanitizeLib

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
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

	// Return cleaned phone number
	return cleanPhone
}

func PagingNumber(pageNoStr string, noPerPageStr string) entity.PagingNumbers {
	// Default values
	pageNo := constant.PAGE_NO
	noPerPage := constant.NO_PER_PAGE

	// Parse pageNoStr to int if it's a valid positive integer
	if pageNoInt, err := strconv.Atoi(pageNoStr); err == nil && pageNoInt > 0 {
		pageNo = pageNoInt
	}

	// Parse noPerPageStr to int if it's a valid positive integer
	if noPerPageInt, err := strconv.Atoi(noPerPageStr); err == nil && noPerPageInt > 0 {
		noPerPage = noPerPageInt
	}

	return entity.PagingNumbers{PageNo: pageNo, NoPerPage: noPerPage}
}
