.PHONY: build run test clean docker-build docker-run docker-stop docker-logs db-create db-migrate lint fmt security docs build-prod

# Build the application
build:
	go build -o bin/shark-saas-backend main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run with hot reload (requires air)
dev:
	air

# Build Docker image
docker-build:
	docker build -t shark-saas-backend .

# Run with Docker Compose
docker-run:
	docker-compose up -d

# Stop Docker containers
docker-stop:
	docker-compose down

# View logs
docker-logs:
	docker-compose logs -f

# Database migration (manual)
db-create:
	mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS shark_saas;"

db-migrate:
	go run main.go

# Linting
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Security check
security:
	gosec ./...

# Generate API documentation
docs:
	swag init

# Production build
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/shark-saas-backend main.go
