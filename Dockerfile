# Use an official Go image as the base image
FROM golang:1.22.4 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -o bin/netlog

# Use a minimal base image for the final image
FROM ubuntu:24.04

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY .env ./
COPY --from=builder /app/bin/netlog .

# Set environment variable
ENV PORT=3001

# Expose the port the app runs on
EXPOSE 3001

# Command to run the application
CMD ["./netlog"]
