# Use an official Go runtime as the base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app (main.go is inside cmd folder)
RUN go build -o /app/main cmd/main.go

# Start a new image from scratch
FROM debian:bullseye-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the 'builder' image
COPY --from=builder /app/main .

# Expose port
EXPOSE 8259

# Command to run the executable
CMD ["./main"]
