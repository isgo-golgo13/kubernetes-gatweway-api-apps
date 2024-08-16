# Variables
APP_NAME := go-chi-rest-app
BUILD_DIR := build
GO_FILES := $(shell find . -name '*.go')
DOCKER_COMPOSE_FILE := docker-compose.yml

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build: $(GO_FILES)
	@echo "Building the Go application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./...
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Run the Go application locally (outside of Docker)
.PHONY: run
run: build
	@echo "Running the Go application..."
	@$(BUILD_DIR)/$(APP_NAME)

# Run the Docker Compose setup
.PHONY: up
up:
	@echo "Starting Docker Compose services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop the Docker Compose setup
.PHONY: down
down:
	@echo "Stopping Docker Compose services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Clean up build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# View logs from the Go application container
.PHONY: logs
logs:
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f app

# Fully clean the environment
.PHONY: full-clean
full-clean: clean
	@echo "Removing Docker volumes and images..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down -v --rmi all --remove-orphans
	@docker system prune -f
	@echo "Full clean complete."
