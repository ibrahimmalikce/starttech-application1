# Stage 1: Build
FROM golang:1.25rc1-alpine AS builder

WORKDIR /app

# Copy go mod files first for better caching
COPY Server/MuchToDo/go.mod Server/MuchToDo/go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY Server/MuchToDo/ .
COPY Server/MuchToDo/.env .env

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Stage 2: Run
FROM alpine:3.18

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

RUN chown appuser:appgroup main

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

CMD ["./main"]
