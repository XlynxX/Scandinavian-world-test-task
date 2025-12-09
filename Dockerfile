# Use the official Golang image as the base
FROM golang:1.25

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=windows \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /app
# Copy the entire application source
COPY . .

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

EXPOSE 34115

# Run the application
CMD ["wails", "dev"]