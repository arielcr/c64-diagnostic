package request

import (
	"github.com/arielcr/c64-diagnostic/model"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateDiagnostic(diag model.Diagnostic) error {
	if err := validate.Struct(diag); err != nil {
		return err
	}
	return nil
}
