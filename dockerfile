FROM golang:1.22.2-alpine AS builder
RUN apk --update add alpine-sdk

# Set working directory
WORKDIR /app

# Copy your Go application code
COPY . .

# Download dependencies after copying the application code
RUN echo $GOPATH
RUN go get -u -d github.com/golang-migrate/migrate github.com/lib/pq
RUN go mod download

# Set environment variables (assuming env.sh sets them)
ENV $(cat /app/commands/env.sh)

# Build the application (adjust the command for your specific needs)
RUN go build -o main .

# Switch to a minimal runtime image
FROM alpine:latest

# Copy the application binary
COPY --from=builder /app/main /catssocial_server

# Expose port 8080
EXPOSE 8080

# Set the entrypoint to run your application
ENTRYPOINT ["/catssocial_server"]