# Makefile

# Build the Docker image
build:
	docker-compose build

# Run the Docker containers
up:
	docker-compose up -d

# Stop the Docker containers
down:
	docker-compose down

# View logs from the app service
logs:
	docker-compose logs -f app

# Clean up Docker environment
clean:
	docker-compose down -v --rmi all --remove-orphans
	docker system prune -f
