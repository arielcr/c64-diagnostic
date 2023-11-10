package request

import (
	"encoding/json"
	"net/http"

	"github.com/arielcr/c64-diagnostic/model"
)

func DecodeDiagnostic(diag *model.Diagnostic, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&diag); err != nil {
		return err
	}
	return nil
}
