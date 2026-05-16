FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /batch-worker ./services/batch-worker/main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /batch-worker /batch-worker
CMD ["/batch-worker"]
