CREATE TABLE pipeline_records (
    id TEXT PRIMARY KEY,

    name TEXT NOT NULL,

    pipeline_job_id TEXT NOT NULL,


    email TEXT,

    phone TEXT,

    status TEXT NOT NULL,

    error_message TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_pipeline_job
    FOREIGN KEY (pipeline_job_id)
    REFERENCES pipeline_jobs(id)
    ON DELETE CASCADE
);