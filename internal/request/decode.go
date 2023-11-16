package request

import (
	"encoding/json"
	"net/http"

	"github.com/arielcr/c64-diagnostic/internal/diagnostics"
)

func DecodeDiagnostic(diag *diagnostics.Diagnostic, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&diag); err != nil {
		return err
	}
	return nil
}

func DecodeDiagnosticStatus(status *diagnostics.DiagnosticStatus, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&status); err != nil {
		return err
	}
	return nil
}
