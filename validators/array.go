package validators

import (
	"github.com/go-playground/validator/v10"
)

func validateArrayRequired(fl validator.FieldLevel) bool {
	return fl.Field().Len() != 0
}
