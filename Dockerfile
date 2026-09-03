#Build the binary
FROM golang:1.25.13-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

#Run the binary
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

RUN adduser -D appuser
USER appuser

EXPOSE 8080
CMD ["./server"]