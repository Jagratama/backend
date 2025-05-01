# Stage 1 - Build
FROM golang:1.24.2-alpine AS builder

# Set the working directory
WORKDIR /build

# Install necessary packages
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application
# RUN go build -o ./api
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2 - Create the final image
FROM gcr.io/distroless/base-debian12

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /build/main .

# Copy the config file
COPY --from=builder /build/.env .env

# Run the binary
CMD ["./main"]