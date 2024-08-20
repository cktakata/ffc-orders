# Use the official Go image from the Docker Hub
FROM golang:1.23 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./src

# Start from a small image to reduce size
FROM debian:bullseye-slim

# Install required packages
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

# Copy the binary to the production image from the build stage
COPY --from=build /app/main /app/main

# Expose port 8080 for the Go app
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["/app/main"]
