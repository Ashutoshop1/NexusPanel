.PHONY: all build run test clean docker help install-tools

# Variables
BINARY_NAME=nexuspanel
SERVER_BINARY=nexuspanel-server
AGENT_BINARY=nexuspanel-agent
GO=go
GOFLAGS=-v
LDFLAGS=-ldflags "-s -w"

all: build

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

install-tools: ## Install development tools
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install github.com/swaggo/swag/cmd/swag@latest

deps: ## Download dependencies
	$(GO) mod download
	$(GO) mod tidy

build: ## Build the application
	@echo "Building server..."
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o bin/$(SERVER_BINARY) ./cmd/server
	@echo "Building agent..."
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o bin/$(AGENT_BINARY) ./cmd/agent

build-server: ## Build only the server
	@echo "Building server..."
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o bin/$(SERVER_BINARY) ./cmd/server

build-agent: ## Build only the agent
	@echo "Building agent..."
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o bin/$(AGENT_BINARY) ./cmd/agent

run: ## Run the server
	$(GO) run ./cmd/server/main.go

run-agent: ## Run the agent
	$(GO) run ./cmd/agent/main.go

test: ## Run tests
	$(GO) test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Run tests with coverage report
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linters
	golangci-lint run --timeout 5m

fmt: ## Format code
	$(GO) fmt ./...
	gofmt -s -w .

vet: ## Run go vet
	$(GO) vet ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -rf dist/
	@rm -f coverage.out coverage.html
	@rm -f $(SERVER_BINARY) $(AGENT_BINARY)
	@echo "Clean complete"

docker: ## Build Docker image
	docker build -t nexuspanel:latest -f deploy/docker/Dockerfile .

docker-compose-up: ## Start services with docker-compose
	docker-compose -f deploy/docker/docker-compose.yml up -d

docker-compose-down: ## Stop services with docker-compose
	docker-compose -f deploy/docker/docker-compose.yml down

web-install: ## Install frontend dependencies
	cd web && npm install

web-dev: ## Run frontend development server
	cd web && npm run dev

web-build: ## Build frontend for production
	cd web && npm run build

web-preview: ## Preview production build
	cd web && npm run preview

dev: ## Run both backend and frontend in development mode
	@echo "Starting development environment..."
	@make -j2 run web-dev

migrate-up: ## Run database migrations
	$(GO) run cmd/server/main.go migrate up

migrate-down: ## Rollback database migrations
	$(GO) run cmd/server/main.go migrate down

proto: ## Generate protobuf files
	protoc --go_out=. --go-grpc_out=. api/proto/*.proto

swag: ## Generate swagger documentation
	swag init -g cmd/server/main.go -o docs/swagger

.DEFAULT_GOAL := help
