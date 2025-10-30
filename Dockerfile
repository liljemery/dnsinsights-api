# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dnsinsight-api ./cmd/dnsinsight-api/main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/dnsinsight-api /app/dnsinsight-api
ENV APP_PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/dnsinsight-api"]
