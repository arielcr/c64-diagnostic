package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/arielcr/c64-diagnostic/internal/adapter/reader"
	"github.com/arielcr/c64-diagnostic/internal/application"
	"github.com/arielcr/c64-diagnostic/internal/diagnostics"
)

const port = ":8081"

func main() {

	logger := InitializeLogger()

	jsonFilePath, err := getJsonFilePath()
	if err != nil {
		slog.Error("JSON File path invalid", "error", err)
		os.Exit(1)
	}

	slog.Info("JSON file path", "path", jsonFilePath)
	reader := reader.NewJsonReader(jsonFilePath, logger)

	service := diagnostics.NewService(reader, logger)

	api := application.NewAPI(logger, service)

	if err := api.Run(port); err != nil {
		slog.Error("Unable to start the application", "error", err)
		os.Exit(1)
	}

}

func InitializeLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}

func getJsonFilePath() (string, error) {
	relativePath := "data/diagnostic.json"

	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", err
	}

	return absolutePath, nil
}
