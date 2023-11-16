package request

import (
	"github.com/arielcr/c64-diagnostic/internal/diagnostics"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateDiagnostic(diag diagnostics.Diagnostic) error {
	if err := validate.Struct(diag); err != nil {
		return err
	}
	return nil
}

func ValidateDiagnosticStatus(diag diagnostics.DiagnosticStatus) error {
	if err := validate.Struct(diag); err != nil {
		return err
	}
	return nil
}
