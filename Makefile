SERVICE_BINARY = CustomerService

up_build: build_service
	@echo "Stopping docker images (if running...)"
	cd ./docker && docker compose down
	@echo "Building (when required) and starting docker images..."
	cd ./docker && docker compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	cd ./docker/ && docker compose down
	@echo "Done!"

build_service:
	@echo "Building service binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ./docker/${SERVICE_BINARY} ./cmd/api
	@echo "Done!"
