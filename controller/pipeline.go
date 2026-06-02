package controller

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"encoding/json"
	"net/http"
	"time"
)

type Record struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CreatePipeline(w http.ResponseWriter, r *http.Request) {
	var records []Record

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&records); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	addJob := models.PipelineJob{
		ID:         utils.GenerateID(),
		FileName:   "data.csv",
		FileType:   "csv",
		SourcePath: "",
		Status:     "PENDING",
		CreatedAt:  time.Now(),
	}

	if err := repository.CreatePipelineJob(addJob); err != nil {
		http.Error(w, "Failed to create pipeline job", http.StatusInternalServerError)
		return
	}

	repository.UpdatePipelineStatus(addJob.ID, "IN_PROGRESS")

	// Process each record
	for i := range records {
		record := &records[i]

		record.Status = "COMPLETED"
		// Validate phone
		if !utils.IsValidPhone(record.Phone) {
			record.Status = "FAILED"
			record.Error = "phone number must contain exactly 10 digits"
		}

		// Validate email
		if !utils.IsValidEmail(record.Email) {
			record.Status = "FAILED"
			record.Error = "invalid email address"
		}

		job := models.PipelineRecord{
			ID:            utils.GenerateID(),
			PipelineJobId: addJob.ID,
			Name:          record.Name,
			Phone:         record.Phone,
			Email:         record.Email,
			Status:        record.Status,
			CreatedAt:     time.Now(),
		}

		go repository.CreatePipelineRecord(job)
	}

	repository.UpdatePipelineStatus(addJob.ID, "COMPLETED")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": addJob.ID})
}

func GetPipelines(w http.ResponseWriter, r *http.Request) {
	jobs, err := repository.GetPipelineJobs()
	if err != nil {
		http.Error(w, "Failed to retrieve pipeline jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func GetPipelineByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/pipelines/"):]
	job, err := repository.GetPipelineJobByID(id)
	if err != nil {
		http.Error(w, "Pipeline job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func GetAllPipelineReport(w http.ResponseWriter, r *http.Request) {
	reports, err := repository.GetAllPipelineReport()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
