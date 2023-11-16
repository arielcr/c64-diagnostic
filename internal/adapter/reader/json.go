package reader

import (
	"encoding/json"
	"log/slog"

	"github.com/arielcr/c64-diagnostic/internal/diagnostics"

	jsonq "github.com/thedevsaddam/gojsonq/v2"
)

const JSON_FILE_PATH = "./data/diagnostic.json"

type JsonReader struct {
	logger *slog.Logger
	jq     *jsonq.JSONQ
}

func NewJsonReader(logger *slog.Logger) *JsonReader {
	newJsonReader := JsonReader{
		logger: logger,
		jq:     jsonq.New().File(JSON_FILE_PATH),
	}

	return &newJsonReader
}

func (r *JsonReader) Query(status diagnostics.DiagnosticStatus) (diagnostics.Diagnostic, error) {

	currentDiagnostic := "diagnostics." + status.CurrentDiagnostic

	result := r.jq.From(currentDiagnostic).Where("step", "=", status.CurrentStep).First()

	var step diagnostics.Step

	jsonString, err := json.Marshal(result)
	if err != nil {
		r.logger.Error(err.Error())
		return diagnostics.Diagnostic{}, err
	}

	err = json.Unmarshal(jsonString, &step)
	if err != nil {
		r.logger.Error(err.Error())
		return diagnostics.Diagnostic{}, err
	}

	r.logger.Info("LA CATEGORIA DEL STEP", "cat", step.Category)

	return diagnostics.Diagnostic{}, nil

}
