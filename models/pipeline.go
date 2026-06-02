package models

import "time"

type PipelineJob struct {
	ID         string           `json:"id"`
	FileName   string           `json:"file_name"`
	FileType   string           `json:"file_type"`
	SourcePath string           `json:"source_path"`
	Status     string           `json:"status"`
	CreatedAt  time.Time        `json:"created_at"`
	Records    []PipelineRecord `json:"records"`
}
