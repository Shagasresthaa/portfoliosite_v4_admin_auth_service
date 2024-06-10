# Start from the official Golang image.
FROM golang:1.22.3 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy the Go module files first to leverage Docker cache.
COPY go.mod go.sum ./

# Download dependencies. They will be cached if the go.mod and go.sum files have not changed.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go app. Disable CGO and target Linux.
# Ensure the build path './cmd' correctly points to where your main.go file is located.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o adminauth ./cmd

# Use a small Alpine base image.
FROM alpine:latest  

# No need for ca-certificates since there are no outgoing HTTPS requests.
# If needed in future, can uncomment the following line:
# RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container.
WORKDIR /root/

# Copy the built binary file from the builder stage.
COPY --from=builder /app/adminauth .

# Also copy the .env file into the container to set environment variables.
COPY --from=builder /app/.env .

# Expose port 8080 on which your app listens.
EXPOSE 8080

# Command to run when starting the container.
CMD ["./adminauth"]
