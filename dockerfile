# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o discount-engine

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Install certificates (needed for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Copy the built binary and rules file from the builder stage
COPY --from=builder /app/discount-engine .
COPY --from=builder /app/go_task_rules.json .

# Expose the port
EXPOSE 8000

# Command to run the application
CMD ["./discount-engine"]
