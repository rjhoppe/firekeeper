# Firekeeper

A Go-based scheduler for sending scheduled GET and POST requests to your API. Firekeeper is designed for reliability, maintainability, and ease of deployment, leveraging Go's standard library and modern scheduling libraries.

## Description

Firekeeper is a lightweight, container-friendly service that allows you to schedule HTTP GET and POST requests to any API endpoint. It is ideal for automating reminders, notifications, or periodic data fetching and posting. Scheduling is handled using cron expressions, and the project is built with Go 1.23+ for performance and simplicity.

### Dependencies

- Go 1.23 or newer
- Docker (optional, for containerized deployment)
- (Optional) Python and pre-commit for code quality hooks

### Installing

1. **Clone the repository:**
   ```sh
   git clone https://github.com/rjhoppe/firekeeper.git
   cd firekeeper
   ```
2. **Install Go dependencies:**
   ```sh
   go mod download
   ```
3. **(Optional) Install pre-commit hooks:**
   ```sh
   pip install pre-commit
   pre-commit install
   ```

### Executing program

#### Run locally
```sh
go run main.go
```

#### Build and run with Docker
```sh
docker build -t firekeeper .
docker run --rm -p 8081:8081 firekeeper
```

#### Configuration
- Edit `main.go` to adjust scheduled jobs, endpoints, and cron expressions as needed.

## Help

- If you encounter connection errors, ensure your API server is running and accessible from the container or host.
- For Docker networking issues, see the [Docker networking docs](https://docs.docker.com/network/).
- To run all pre-commit hooks manually:
```sh
pre-commit run --all-files
```
