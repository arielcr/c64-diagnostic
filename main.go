package main

import (
	"log/slog"
	"os"

	"github.com/arielcr/c64-diagnostic/app"
)

const port = ":8080"

func main() {
	api := app.NewAPI()

	if err := api.Run(port); err != nil {
		slog.Error("Unable to start the application", "error", err)
		os.Exit(1)
	}

}
