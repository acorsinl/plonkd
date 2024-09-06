GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=plonkd

# Default target executed when no target is specified
all: build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Install dependencies
deps:
	$(GOGET) -u ./...

# Run the application
run:
	./$(BINARY_NAME)

# Docker build and run
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker compose up -d

# Display available targets
help:
	@echo "Available targets:"
	@echo "  all           - Build the project"
	@echo "  build         - Build the project"
	@echo "  clean         - Clean the project"
	@echo "  test          - Run tests"
	@echo "  deps          - Install dependencies"
	@echo "  run           - Run the application"
	@echo "  docker-build  - Build the Docker image"
	@echo "  docker-run    - Run the Docker container"%
