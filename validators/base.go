package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func InitValidator() error {
	validate = validator.New()

	err := validate.RegisterValidation("phone_number", validatePhoneNumber)
	if err != nil {
		return fmt.Errorf("Error register phone number validation { %s } ", err)
	}

	err = validate.RegisterValidation("array_required", validateArrayRequired)
	if err != nil {
		return fmt.Errorf("Error register array required validation { %s } ", err)
	}

	err = validate.RegisterValidation("uuid", validateUuid)
	if err != nil {
		return fmt.Errorf("Error registering uuid validation: %s ", err.Error())
	}

	return nil
}

func ValidateInputs(request interface{}) error {
	err := validate.Struct(request)
	if err == nil {
		return nil
	}

	if err, ok := err.(*validator.InvalidValidationError); ok {
		return fmt.Errorf("Invalid validation error: %s ", err.Error())
	}

	var errorMessages []string
	reflectedRequest := reflect.ValueOf(request)

	for _, errorComponent := range err.(validator.ValidationErrors) {
		field, isFieldFound := reflectedRequest.Type().FieldByName(errorComponent.StructField())
		if !isFieldFound {
			return fmt.Errorf("Found validation error for non-existing field %s ", errorComponent.StructField())
		}
		fieldKey := getStructFieldKey(field, strings.ToLower(errorComponent.StructField()))

		switch errorComponent.Tag() {
		case "required":
			errorMessages = append(errorMessages,
				strings.Join([]string{
					"The", fieldKey, "field is required",
				}, " "),
			)
		case "email":
			errorMessages = append(errorMessages,
				strings.Join([]string{
					"The", fieldKey, "should be a valid email",
				}, " "),
			)
		case "phone_number":
			errorMessages = append(errorMessages,
				strings.Join([]string{
					"The", fieldKey, "should be a valid phone number",
				}, " "),
			)
		case "array_required":
			errorMessages = append(errorMessages,
				strings.Join([]string{
					"The", fieldKey, "field is required",
				}, " "),
			)
		}

	}
	return errors.New(strings.Join(errorMessages, "\n"))
}

func getStructFieldKey(structField reflect.StructField, defaultKey string) string {
	if fieldJsonKey := structField.Tag.Get("json"); len(fieldJsonKey) > 0 {
		return fieldJsonKey
	}

	if fieldSchemaKey := structField.Tag.Get("schema"); len(fieldSchemaKey) > 0 {
		return fieldSchemaKey
	}

	return defaultKey
}
