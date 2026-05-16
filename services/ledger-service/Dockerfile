FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /ledger-service ./services/ledger-service/cmd/api

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /ledger-service /ledger-service
EXPOSE 8080
CMD ["/ledger-service"]
