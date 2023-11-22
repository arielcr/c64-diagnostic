package main

import (
	"log/slog"
	"os"

	"github.com/arielcr/c64-diagnostic/internal/adapter/reader"
	"github.com/arielcr/c64-diagnostic/internal/application"
	"github.com/arielcr/c64-diagnostic/internal/diagnostics"
)

const port = ":8081"

func main() {
	logger := InitializeLogger()

	reader := reader.NewJsonReader(logger)

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
