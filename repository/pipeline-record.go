package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"fmt"
	"time"
)

func CreatePipelineRecord(job models.PipelineRecord) error {
	time.Sleep(10 * time.Second) // Simulate processing time
	query := `
		INSERT INTO pipeline_records
		(id, name, email, phone, status, created_at, error_message, pipeline_job_id)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		job.ID,
		job.Name,
		job.Email,
		job.Phone,
		job.Status,
		job.CreatedAt,
		job.ErrorMessage,
		job.PipelineJobId,
	)

	fmt.Println("Inserted record with ID:", job.ID, "Status:", job.Status, "Error:", job.ErrorMessage, err)

	return err
}
