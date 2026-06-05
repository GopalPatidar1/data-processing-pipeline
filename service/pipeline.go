package service

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
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

func ProcessPipeline(jobID string, records []Record) {
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
			PipelineJobId: jobID,
			Name:          record.Name,
			Phone:         record.Phone,
			Email:         record.Email,
			Status:        record.Status,
			CreatedAt:     time.Now(),
		}

		repository.CreatePipelineRecord(job)
	}
	repository.UpdatePipelineStatus(jobID, "COMPLETED")
}

func CreatePipeline(w http.ResponseWriter, r *http.Request) (string, error) {
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
		return "", err
	}

	return addJob.ID, nil
}

func UpdatePipelineStatus(jobID, status string) error {
	return repository.UpdatePipelineStatus(jobID, status)
}

func GetPipelines() ([]models.PipelineJob, error) {
	return repository.GetPipelineJobs()
}

func GetPipelineByID(id string) (*repository.PipelineJobResponse, error) {
	return repository.GetPipelineJobByID(id)
}

func DeletePipelineById(id string) error {
	return repository.DeletePipelineJobByID(id)
}

func GetAllPipelineReport() ([]models.PipelineJob, error) {
	return repository.GetAllPipelineReport()
}
