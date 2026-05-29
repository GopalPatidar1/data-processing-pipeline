CREATE TABLE pipeline_jobs (
    id TEXT PRIMARY KEY,

    file_name TEXT NOT NULL,

    file_type TEXT NOT NULL,

    source_path TEXT NOT NULL,

    status TEXT NOT NULL DEFAULT 'PENDING',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);