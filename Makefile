# Makefile for building and running the Go application

# Image name
IMAGE_NAME=myapp

# Container name
CONTAINER_NAME=myapp-container

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the application in a Docker container
run: build
	docker run --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

# Run tests in a Docker container
test:
	go test ./...

# Stop the running container
stop:
	-docker stop $(CONTAINER_NAME)

# Remove the running container
remove-container:
	-docker rm $(CONTAINER_NAME)

# Remove the Docker image
clean: stop remove-container
	docker rmi $(IMAGE_NAME)

# Rebuild the application
rebuild: clean build run

# Run all targets
all: build run
