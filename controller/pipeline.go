package controller

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CreatePipeline(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type RequestBody struct {
		FileName   string `json:"file_name"`
		FileType   string `json:"file_type"`
		SourcePath string `json:"source_path"`
	}

	var body RequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Received pipeline job for file:", body.FileName)

	job := models.PipelineJob{
		ID:         utils.GenerateID(),
		FileName:   body.FileName,
		FileType:   body.FileType,
		SourcePath: body.SourcePath,
		Status:     "PENDING",
		CreatedAt:  time.Now(),
	}

	err = repository.CreatePipelineJob(job)

	if err != nil {
		http.Error(w, "Failed to create pipeline job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(job)
}
