.PHONY: migrate-up migrate-down migrate-force develop-run local-run linter psql gen run build test deps-reset \
 tidy deps-upgrade deps-cleancache down-local clean logs-local

APP_NAME = exchange
APP_PORT = 9000
BUILD_DIR = build
MIGRATIONS_FOLDER = internal/repository/migrations
DATABASE_URL = postgres://postgres:password@localhost:5433/auth_db?sslmode=disable
FILES := $(shell docker ps -aq)

# ==============================================================================
# Go migrate postgresql

migrate-up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate-force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)
# ==============================================================================
# Docker compose commands

develop-run:
	echo "Starting docker environment"
	docker-compose -f deployments/docker-compose.dev.yml up --build

local-run:
	echo "Starting local environment"
	docker-compose -f deployments/docker-compose.local.yml up --build

# ==============================================================================
# Tools commands

linter:
	echo "Starting linters"
	golangci-lint run ./...
	gosec -tests ./...

psql:
	docker exec -it $(APP_NAME)-postges psql -U postgres -d password

gen:
	go generate ./..

# ==============================================================================
# Main

run: build
	$(BUILD_DIR)/$(APP_NAME)

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) cmd/$(APP_NAME)/main.go

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover $(go list ./... | grep -v /pgdata/)
	go tool cover -func=cover.out

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Docker support

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)
