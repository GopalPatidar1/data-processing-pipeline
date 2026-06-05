package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"time"
)

type RecordSummary struct {
	Completed  int `json:"completed"`
	Failed     int `json:"failed"`
	InProgress int `json:"inProgress"`
}
type PipelineJobResponse struct {
	ID         string        `json:"id"`
	FileName   string        `json:"file_name"`
	FileType   string        `json:"file_type"`
	SourcePath string        `json:"source_path"`
	Status     string        `json:"status"`
	CreatedAt  time.Time     `json:"created_at"`
	Records    RecordSummary `json:"records"`
}

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

func GetPipelineJobByID(id string) (*PipelineJobResponse, error) {

	query := `
		SELECT
			j.id,
			j.file_name,
			j.file_type,
			j.source_path,
			j.status,
			j.created_at,
			COUNT(*) FILTER (WHERE r.status = 'COMPLETED') AS completed,
			COUNT(*) FILTER (WHERE r.status = 'FAILED') AS failed,
			COUNT(*) FILTER (WHERE r.status = 'IN_PROGRESS') AS in_progress
		FROM pipeline_jobs j
		LEFT JOIN pipeline_records r
			ON r.pipeline_job_id = j.id
		WHERE j.id = $1
		GROUP BY
			j.id,
			j.file_name,
			j.file_type,
			j.source_path,
			j.status,
			j.created_at
	`

	var job PipelineJobResponse

	err := config.DB.QueryRow(context.Background(), query, id).Scan(
		&job.ID,
		&job.FileName,
		&job.FileType,
		&job.SourcePath,
		&job.Status,
		&job.CreatedAt,
		&job.Records.Completed,
		&job.Records.Failed,
		&job.Records.InProgress,
	)

	if err != nil {
		return nil, err
	}

	return &job, nil
}

func GetAllPipelineReport() ([]models.PipelineJob, error) {

	query := `
		   SELECT
           j.id,
           j.file_name,
           j.file_type,
           j.source_path,
           j.status,
           j.created_at,
           COALESCE(
               JSON_AGG(
                   JSON_BUILD_OBJECT(
                       'id', r.id,
                       'name', r.name,
                       'email', r.email,
                       'phone', r.phone,
                       'status', r.status,
					   'pipeline_job_id', r.pipeline_job_id
                   )
               ) FILTER (WHERE r.id IS NOT NULL),
               '[]'::json
           ) AS records
       FROM pipeline_jobs j
       LEFT JOIN pipeline_records r
           ON r.pipeline_job_id = j.id
       GROUP BY
           j.id;
	`

	rows, err := config.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reports []models.PipelineJob

	for rows.Next() {

		var report models.PipelineJob

		err := rows.Scan(
			&report.ID,
			&report.FileName,
			&report.FileType,
			&report.SourcePath,
			&report.Status,
			&report.CreatedAt,
			&report.Records,
		)

		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}
func DeletePipelineJobByID(id string) error {

	tx, err := config.DB.Begin(context.Background())
	if err != nil {
		return err
	}

	// Rollback if anything fails
	defer tx.Rollback(context.Background())

	// Delete pipeline records
	recordsQuery := `
		DELETE FROM pipeline_records
		WHERE pipeline_job_id = $1
	`

	_, err = tx.Exec(
		context.Background(),
		recordsQuery,
		id,
	)

	if err != nil {
		return err
	}

	// Delete pipeline job
	jobQuery := `
		DELETE FROM pipeline_jobs
		WHERE id = $1
	`

	_, err = tx.Exec(
		context.Background(),
		jobQuery,
		id,
	)

	if err != nil {
		return err
	}

	// Commit transaction
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
