# Data Processing Pipeline API

A simple REST API built with Go for managing pipeline jobs and records.

## Project Structure

```text
.
├── cmd/
│   └── main.go
├── controller/
├── service/
├── repository/
├── routes/
├── models/
├── config/
├── utils/
└── tests/
```

### Layers

* **Routes** – Registers API endpoints.
* **Controller** – Handles HTTP requests and responses.
* **Service** – Contains business logic.
* **Repository** – Handles database operations.
* **Models** – Defines data structures.
* **Config** – Database configuration and application setup.
* **Utils** – Utility/helper functions.

---

## Prerequisites

* Go 1.22+
* PostgreSQL

---

## Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE pipeline_db;
```

Update the connection string in the configuration file:

```go
postgres://username:password@localhost:5432/pipeline_db
```

---

## Running the Application

Start the API server:

```bash
go run cmd/main.go
```

Server starts on:

```text
http://localhost:3007
```

---

## Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run tests for a specific package:

```bash
go test ./service
```

---

## API Endpoints

### Create Pipeline Job

**POST**

```http
/api/v1/pipelines
```

#### Request Body

```json
[
  {
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "9876543210"
  }
]
```

#### Response

```json
{
  "id": "pipeline-job-id"
}
```

---

### Get All Pipeline Jobs

**GET**

```http
/api/v1/pipelines
```

Returns all pipeline jobs.

---

### Get All Pipeline Jobs With Records

**GET**

```http
/api/v1/pipelines/all
```

Returns all pipeline jobs along with aggregated pipeline records.

#### Response

```json
[
  {
    "id": "job-id",
    "file_name": "data.csv",
    "status": "COMPLETED",
    "records": [
      {
        "id": "record-id",
        "name": "John Doe",
        "email": "john@example.com",
        "status": "COMPLETED"
      }
    ]
  }
]
```

---

### Get Pipeline By ID

**GET**

```http
/api/v1/pipelines/{id}
```

Returns a single pipeline job with its associated records.

---

### Delete Pipeline By ID

**DELETE**

```http
/api/v1/pipelines/{id}
```

Deletes a pipeline job and its associated records.

#### Response

```http
204 No Content
```

---

## Processing Flow

1. Create a pipeline job.
2. Pipeline status is set to `IN_PROGRESS`.
3. Records are validated:

   * Email validation
   * Phone validation
4. Records are stored in the database.
5. Pipeline status is updated to `COMPLETED`.

---

## Features

* RESTful API
* Layered architecture
* Background job processing using Goroutines
* PostgreSQL integration using pgx
* Email and phone validation
* Aggregated pipeline reporting
* Unit test support
* Separation of controller, service, and repository layers

---

## Author

Gopal Patidar
