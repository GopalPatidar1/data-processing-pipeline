package routes

import (
	"backend/controller"
	"net/http"
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
}
