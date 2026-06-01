package repository

import (
	"backend/config"
	"backend/models"
	"context"
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

func GetPipelineJobs() ([]models.PipelineJob, error) {
	query := `
		SELECT
			id,
			file_name,
			file_type,
			source_path,
			status,
			created_at
		FROM pipeline_jobs
	`

	rows, err := config.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.PipelineJob

	for rows.Next() {

		var job models.PipelineJob

		err := rows.Scan(
			&job.ID,
			&job.FileName,
			&job.FileType,
			&job.SourcePath,
			&job.Status,
			&job.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func UpdatePipelineStatus(id string, status string) error {

	query := `
		UPDATE pipeline_jobs
		SET status = $1
		WHERE id = $2
	`

	_, err := config.DB.Exec(
		context.Background(),
		query,
		status,
		id,
	)

	return err
}
