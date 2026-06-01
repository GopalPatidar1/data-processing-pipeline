package repository

import (
	"backend/config"
	"backend/models"
	"context"
)

func CreatePipelineRecord(job models.PipelineRecord) error {
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

	return err
}
