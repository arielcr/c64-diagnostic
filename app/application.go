package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/arielcr/c64-diagnostic/model"
	"github.com/arielcr/c64-diagnostic/request"
	"github.com/arielcr/c64-diagnostic/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type API struct {
	Router *mux.Router
	Log    *slog.Logger
}

func NewAPI() *API {
	s := API{
		Router: mux.NewRouter(),
	}
	return &s
}

func (a *API) InitializeRoutes() {
	a.Router.HandleFunc("/diagnose", a.Diagnose).Methods("POST")
}

func (a *API) InitializeLogger() {
	a.Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(a.Log)
}

func (a *API) Run(addr string) error {
	a.InitializeLogger()
	a.InitializeRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://c64diagnostic.com", "http://127.0.0.1:5500"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(a.Router)

	a.Log.Info("c64-diagnostic service started", "port", addr)

	return http.ListenAndServe(addr, handler)
}

func (a *API) Diagnose(w http.ResponseWriter, r *http.Request) {
	var diag model.Diagnostic

	if err := request.DecodeDiagnostic(&diag, r); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := request.ValidateDiagnostic(diag); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request structure or missing parameters")
		return
	}

	defer r.Body.Close()

	slog.Info("Diagnostic processed", "payload", diag)

	utils.RespondWithJSON(w, http.StatusOK, diag)

}
