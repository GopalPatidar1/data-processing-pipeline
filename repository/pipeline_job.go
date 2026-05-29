package repository

import (
	"context"

	"backend/config"
	"backend/models"
)

func CreatePipelineJob(job models.PipelineJob) error {
	query := `
		INSERT INTO pipeline_jobs
		(id, file_name, file_type, source_path, status, created_at)
		VALUES
		($1, $2, $3, $4, $5, $6)
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		job.ID,
		job.FileName,
		job.FileType,
		job.SourcePath,
		job.Status,
		job.CreatedAt,
	)

	return err
}
