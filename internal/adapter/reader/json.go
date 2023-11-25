package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/arielcr/c64-diagnostic/internal/diagnostics"

	jsonq "github.com/thedevsaddam/gojsonq/v2"
)

const jsonFilePath = "diagnostic.json"

type JsonReader struct {
	logger *slog.Logger
	jq     *jsonq.JSONQ
}

func NewJsonReader(logger *slog.Logger) *JsonReader {
	newJsonReader := JsonReader{
		logger: logger,
	}

	return &newJsonReader
}

func (r *JsonReader) SetDataSource(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	jsonString := string(body)

	r.jq = jsonq.New().FromString(jsonString)

	return nil
}

func (r *JsonReader) GetStep(diagnostic string, stepNumber string) (diagnostics.Step, error) {
	result := r.jq.Reset().From("diagnostics."+diagnostic).Where("step", "=", stepNumber).First()

	step, err := r.ParseStepFromJSON(result)
	if err != nil {
		return diagnostics.Step{}, err
	}

	return step, nil
}

func (r *JsonReader) GetDiagnosticMeta() (diagnostics.DiagnosticMeta, error) {
	result := r.jq.Reset().From("diagnostic_meta").Get()

	meta, err := r.ParseMetaFromJSON(result)
	if err != nil {
		return diagnostics.DiagnosticMeta{}, err
	}

	return meta, nil
}

func (r *JsonReader) ParseStepFromJSON(data interface{}) (diagnostics.Step, error) {
	var step diagnostics.Step

	jsonString, err := json.Marshal(data)
	if err != nil {
		return diagnostics.Step{}, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	err = json.Unmarshal(jsonString, &step)
	if err != nil {
		return diagnostics.Step{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return step, nil
}

func (r *JsonReader) ParseMetaFromJSON(data interface{}) (diagnostics.DiagnosticMeta, error) {
	var step diagnostics.DiagnosticMeta

	jsonString, err := json.Marshal(data)
	if err != nil {
		return diagnostics.DiagnosticMeta{}, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	err = json.Unmarshal(jsonString, &step)
	if err != nil {
		return diagnostics.DiagnosticMeta{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return step, nil
}
