package controller

import (
	"backend/service"
	"encoding/json"
	"net/http"
)

func CreatePipeline(w http.ResponseWriter, r *http.Request) {
	var records []service.Record

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&records); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	jobId, err := service.CreatePipeline(w, r)
	if err != nil {
		return
	}

	service.UpdatePipelineStatus(jobId, "IN_PROGRESS")

	go service.ProcessPipeline(jobId, records)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": jobId})
}

func GetPipelines(w http.ResponseWriter, r *http.Request) {
	jobs, err := service.GetPipelines()
	if err != nil {
		http.Error(w, "Failed to retrieve pipeline jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func GetPipelineByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/pipelines/"):]
	job, err := service.GetPipelineByID(id)
	if err != nil {
		http.Error(w, "Pipeline job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func DeletePipelineById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/pipelines/"):]
	err := service.DeletePipelineById(id)
	if err != nil {
		http.Error(w, "Pipeline job not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllPipelineReport(w http.ResponseWriter, r *http.Request) {
	reports, err := service.GetAllPipelineReport()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
