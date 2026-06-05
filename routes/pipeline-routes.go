package routes

import (
	"backend/controller"
	"net/http"
	"strings"
)

func RegisterPipelineRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/v1/pipelines", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			controller.CreatePipeline(w, r)

		case http.MethodGet:
			controller.GetPipelines(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Pipeline by ID & Pipeline Report
	mux.HandleFunc("/api/v1/pipelines/", func(w http.ResponseWriter, r *http.Request) {

		switch {

		case r.Method == http.MethodGet &&
			strings.HasSuffix(r.URL.Path, "/all"):

			controller.GetAllPipelineReport(w, r)

		// GET /api/v1/pipelines/{id}
		case r.Method == http.MethodGet:
			controller.GetPipelineByID(w, r)

		// DELETE /api/v1/pipelines/{id}
		case r.Method == http.MethodDelete:
			controller.DeletePipelineById(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
