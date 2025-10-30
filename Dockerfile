# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder
WORKDIR /app

ENV GOTOOLCHAIN=auto

COPY go.mod .
RUN go mod download

COPY . .
# Install swag and generate OpenAPI docs (compiled into the binary)
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN /go/bin/swag init -g routes/router.go -o ./docs

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dnsinsight-api ./cmd/dnsinsight-api/main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/dnsinsight-api /app/dnsinsight-api
ENV APP_PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/dnsinsight-api"]
