package models

import "time"

type PipelineRecord struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	PipelineJobId string    `json:"pipeline_job_id"`
	Status        string    `json:"status"`
	Phone         string    `json:"phone"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}
