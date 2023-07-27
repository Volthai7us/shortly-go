# Start from golang base image
FROM golang:1.17-alpine as builder

# Set the current Working Directory inside the container
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the application
RUN go build -o main ./cmd/url-shortener

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary file from builder
COPY --from=builder /app/main .

# Expose port 5173 to the outside
EXPOSE 5173

# Run the executable
CMD ["./main"]
