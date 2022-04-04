package validators

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

func validatePhoneNumber(fl validator.FieldLevel) bool {
	normalizedPhoneNumber, err := normalizePhoneNumber(fl.Field().String())
	if err != nil {
		return false
	}

	_, i := utf8.DecodeRuneInString(normalizedPhoneNumber)
	prefixTrimmedPhoneNumber := normalizedPhoneNumber[i+2:]

	validCharsRegex := regexp.MustCompile("[0-9]")
	phoneNumberWithoutValidChars := validCharsRegex.ReplaceAllString(prefixTrimmedPhoneNumber, "")

	if len(phoneNumberWithoutValidChars) > 0 {
		log.Printf("Phone number %s contains invalid non-numeric characters", normalizedPhoneNumber)
		return false
	}

	length := len(normalizedPhoneNumber)
	if length < 8 || length > 15 {
		log.Printf("Invalid phone number of length %v", length)
		return false
	}

	return true
}

func normalizePhoneNumber(phoneNumber string) (string, error) {
	whitespaceTrimmedPhoneNumber := strings.TrimSpace(phoneNumber)

	prefixNormalizedPhoneNumber, err := normalizePhoneNumberPrefix(whitespaceTrimmedPhoneNumber)
	if err != nil {
		return whitespaceTrimmedPhoneNumber, err
	}

	return prefixNormalizedPhoneNumber, nil
}

func normalizePhoneNumberPrefix(phoneNumber string) (string, error) {
	countryCodePrefix := "+62"

	if strings.HasPrefix(phoneNumber, countryCodePrefix) {
		return phoneNumber, nil
	}

	if !strings.HasPrefix(phoneNumber, "0") {
		return phoneNumber, fmt.Errorf("Invalid prefix for phone number %s ", phoneNumber)
	}

	_, i := utf8.DecodeRuneInString(phoneNumber)
	prefixTrimmedPhoneNumber := phoneNumber[i:]
	return strings.Join([]string{countryCodePrefix, prefixTrimmedPhoneNumber}, ""), nil
}
