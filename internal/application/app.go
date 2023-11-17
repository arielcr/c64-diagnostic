package application

import (
	"log/slog"
	"net/http"

	"github.com/arielcr/c64-diagnostic/internal/diagnostics"
	"github.com/arielcr/c64-diagnostic/internal/request"
	"github.com/arielcr/c64-diagnostic/internal/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type API struct {
	Router  *mux.Router
	Log     *slog.Logger
	Service *diagnostics.Service
}

func NewAPI(logger *slog.Logger, service *diagnostics.Service) *API {
	api := API{
		Router:  mux.NewRouter(),
		Log:     logger,
		Service: service,
	}
	return &api
}

func (a *API) InitializeRoutes() {
	a.Router.HandleFunc("/diagnose", a.Diagnose).Methods("POST")
}

func (a *API) Run(addr string) error {
	a.InitializeRoutes()

	a.Log.Info("c64-diagnostic service started", "port", addr)

	return http.ListenAndServe(addr, a.getAPIHandler())
}

func (a *API) getAPIHandler() http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://c64diagnostic.com", "http://127.0.0.1:5500"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
	})

	return c.Handler(a.Router)
}

func (a *API) Diagnose(w http.ResponseWriter, r *http.Request) {
	var status diagnostics.DiagnosticStatus

	if err := request.DecodeDiagnosticStatus(&status, r); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := request.ValidateDiagnosticStatus(status); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request structure or missing parameters")
		return
	}

	defer r.Body.Close()

	result, err := a.Service.GetNextStep(status)

	if err != nil {
		a.Log.Error("error parsing the diagnostic", "error", err.Error())
		utils.RespondWithError(w, http.StatusBadRequest, "error parsing the diagnostic")
	}

	slog.Info("Diagnostic processed", "payload", result)

	utils.RespondWithJSON(w, http.StatusOK, result)

}
