# 🎓 Student API

A clean, modular RESTful CRUD API built in Go for managing student records. Designed following clean architecture principles, utilizing the standard library router (`net/http`), zero-CGO SQLite for persistence, and structured logging with `slog`.

---

## 🚀 Features

- **Standard Library Routing**: Uses Go 1.22+ enhanced `net/http.ServeMux` for pattern matching and path values.
- **Clean Architecture**: Decoupled layers with storage interfaces, HTTP handlers, centralized types, and standardized responses.
- **Zero-CGO SQLite**: Powered by `modernc.org/sqlite`—compiles and runs cross-platform without requiring CGO or an external C compiler.
- **Robust Configuration**: Managed by `cleanenv`, supporting both YAML configuration files and environment variable overrides.
- **Request Validation**: Automatic input validation and user-friendly error formatting using `go-playground/validator`.
- **Structured Logging**: Built-in structured logging with `log/slog`.
- **Graceful Shutdown**: Intercepts `SIGINT` / `SIGTERM` signals to cleanly drain connections and close database handles with a configurable timeout.

---

## 📁 Project Structure

```text
.
├── cmd/
│   ├── config/
│   │   └── local.yaml              # Local configuration file (fallback)
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go           # Configuration loader (cleanenv)
│   │   ├── http/
│   │   │   └── handler/
│   │   │       └── student/
│   │   │           └── student.go  # HTTP handlers for Student CRUD
│   │   ├── storage/
│   │   │   ├── sqlite/
│   │   │   │   └── sqlite.go       # SQLite implementation
│   │   │   └── storage.go          # Storage interface definition
│   │   ├── types/
│   │   │   └── types.go            # Data structures & validation tags
│   │   └── utils/
│   │       └── response/
│   │           └── response.go     # Standardized JSON response helpers
│   ├── storage/
│   │   └── storage.db              # SQLite database file
│   └── student-api/
│       └── main.go                 # Application entry point
├── config/
│   └── local.yaml                  # Root configuration file
├── go.mod
├── go.sum
└── README.md
```

---

## 🛠️ Getting Started

### Prerequisites

- **Go**: Version `1.22` or later installed ([Download Go](https://go.dev/dl/)).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Gautam-aman/Student-Api.git
   cd Student-Api
   ```

2. Download Go dependencies:
   ```bash
   go mod download
   ```

---

## ⚙️ Configuration

Application settings are managed in `config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"

http_server:
  address: "localhost:3000"
```

You can pass the configuration path via:
- **CLI Flag (Recommended)**: `--config <path>` or `-config <path>`
- **Environment Variable**: `CONFIG_PATH=<path>`

---

## ▶️ Running the Application

Start the server using the configuration file:

```bash
go run cmd/student-api/main.go -config config/local.yaml
```

Or using an environment variable:

```bash
CONFIG_PATH=config/local.yaml go run cmd/student-api/main.go
```

Expected output:
```text
2026/11/11 15:15:18 INFO storage initialized env=dev version=1.0.0
2026/11/11 15:15:18 INFO server started address=localhost:3000
```

To stop the server safely, press `Ctrl + C`. The server will execute a graceful shutdown.

---

## 📡 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/api/students` | Create a new student |
| `GET` | `/api/students` | Retrieve all students |
| `GET` | `/api/students/{id}` | Retrieve a student by ID |

---

### 1. Create a Student
**Request:**
- **URL**: `/api/students`
- **Method**: `POST`
- **Headers**: `Content-Type: application/json`
- **Body**:
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com",
    "age": 22
  }
  ```

**Response (`201 Created`):**
```json
{
  "id": 1
}
```

**Validation Error (`400 Bad Request`):**
```json
{
  "status": "Error",
  "error": "field Name is a required field"
}
```

---

### 2. Get All Students
**Request:**
- **URL**: `/api/students`
- **Method**: `GET`

**Response (`200 OK`):**
```json
[
  {
    "id": 1,
    "name": "Jane Doe",
    "email": "jane@example.com",
    "age": 22
  }
]
```
*(Returns `[]` when no records exist).*

---

### 3. Get Student by ID
**Request:**
- **URL**: `/api/students/1`
- **Method**: `GET`

**Response (`200 OK`):**
```json
{
  "id": 1,
  "name": "Jane Doe",
  "email": "jane@example.com",
  "age": 22
}
```

**Not Found (`404 Not Found`):**
```json
{
  "status": "Error",
  "error": "student not found with id: 99"
}
```

---

## 🧪 Testing with cURL

Here are quick commands to test your API in the terminal:

```bash
# 1. Create student
curl -X POST http://localhost:3000/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Aman","email":"aman@example.com","age":21}'

# 2. Get all students
curl http://localhost:3000/api/students

# 3. Get student by ID
curl http://localhost:3000/api/students/1

# 4. Test 404 (non-existent ID)
curl http://localhost:3000/api/students/999
```

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).

