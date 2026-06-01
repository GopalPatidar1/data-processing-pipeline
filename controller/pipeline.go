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

		// Validate phone
		if !utils.IsValidPhone(record.Phone) {
			record.Status = "FAILED"
			record.Error = "phone number must contain exactly 10 digits"
			continue
		}

		// Validate email
		if !utils.IsValidEmail(record.Email) {
			record.Status = "FAILED"
			record.Error = "invalid email address"
			continue
		}

		job := models.PipelineRecord{
			ID:            utils.GenerateID(),
			PipelineJobId: addJob.ID,
			Name:          record.Name,
			Phone:         record.Phone,
			Email:         record.Email,
			Status:        "PENDING",
			CreatedAt:     time.Now(),
		}

		if err := repository.CreatePipelineRecord(job); err != nil {
			record.Status = "FAILED"
			record.Error = "failed to create job"
			continue
		}

		record.Status = "SUCCESS"
		record.Error = ""
	}

	repository.UpdatePipelineStatus(addJob.ID, "COMPLETED")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(records)
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
