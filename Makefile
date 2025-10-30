APP_NAME=dnsinsight-api
MAIN=./cmd/dnsinsight-api/main.go
SWAG_BIN=swag

.PHONY: run build test tidy docs migrate deps

run:
	go run $(MAIN)

build:
	go build -o bin/$(APP_NAME) $(MAIN)

test:
	go test ./...

tidy:
	go mod tidy

# Generate Swagger docs based on main router file
# Requires: go install github.com/swaggo/swag/cmd/swag@latest
# Output: ./docs
# The router entry file lives at routes/router.go
# Keep this in sync if router file moves

docs:
	$(SWAG_BIN) init -g routes/router.go -o ./docs

# Run DB migrations
migrate:
	go run ./cmd/migrate/main.go

# Install common developer tools locally
deps:
	go install github.com/swaggo/swag/cmd/swag@latest
