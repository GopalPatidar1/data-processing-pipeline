package routes

import (
	"backend/controller"
	"net/http"
)

func RegisterPipelineRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/v1/pipelines", controller.CreatePipeline)
}
