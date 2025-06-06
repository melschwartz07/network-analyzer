# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /network-analyzer

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /network-analyzer /app/network-analyzer
EXPOSE 8080
CMD ["/app/network-analyzer"]