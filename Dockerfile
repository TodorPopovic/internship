# Use the official Golang image to build the application
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Generate wire dependencies
RUN go generate ./cmd

# Set the working directory to the cmd directory
WORKDIR /app/cmd

# Build the Go app
RUN go build -o /app/main .

# Use a minimal Debian-based image for the final stage
FROM debian:stable-slim

# Install ca-certificates (needed for HTTPS requests)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main /usr/local/bin/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/usr/local/bin/main"]
