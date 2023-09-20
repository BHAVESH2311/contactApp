package validators

import (
	"regexp"
	"strings"
)

func ValidateName(value string) bool {

	if value != "" && regexp.MustCompile("^[a-zA-Z ]{2,30}$").MatchString(value) {
		return true
	}
	return false
}

func ValidateContactDetails(typeOfContact string, contactInfo string) (bool, string) {
	if strings.EqualFold(typeOfContact, "phone") || strings.EqualFold(typeOfContact, "email") {
		if strings.EqualFold(typeOfContact, "phone") {
			if regexp.MustCompile("^[0-9]{10}$").MatchString(contactInfo) {
				return true, "Success"
			}
			return false, "Invalid phone number"
		}

		if regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(contactInfo) {
			return true, "Success"
		}
		return false, "Invalid email Id"
	}
	return false, "Invalid type of contact. Please enter 'phone' or 'email'"
}
