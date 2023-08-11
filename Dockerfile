# Use the official Go image as a base image
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o myapp

# Use a minimal image for the runtime
FROM gcr.io/distroless/base-debian10

# Copy the binary from the builder stage
COPY --from=builder /app/myapp /myapp

# Command to run the application~
CMD ["/myapp"]
