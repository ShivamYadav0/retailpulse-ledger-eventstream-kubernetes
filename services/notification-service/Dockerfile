FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /notification-service ./services/notification-service/main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /notification-service /notification-service
CMD ["/notification-service"]
