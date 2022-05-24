### Required tools
### Required tools
GOTOOLS_CHECK = dep gin golangci-lint

all: check_tools fmt ensure-deps test linter

### Tools & dependencies
check_tools:
	@# https://stackoverflow.com/a/25668869
	@echo "Found tools: $(foreach tool,$(GOTOOLS_CHECK),\
        $(if $(shell which $(tool)),$(tool),$(error "No $(tool) in PATH")))"

### Testing
test:
	export GO111MODULE=on; \
	go test ./... -covermode=atomic -coverpkg=./... -count=1 -race

test-cover:
	go test ./... -covermode=atomic -coverprofile=/tmp/coverage.out -coverpkg=./... -count=1
	go tool cover -html=/tmp/coverage.out

### Formatting, linting, and deps
fmt:
	go fmt ./...

linter:
	@echo "==> Running linter..."
	golangci-lint run ./... --fix

dep-init-go-12:
	@echo "==> Initializing project for GO ..."
	export GO111MODULE=on; \
	go mod init

dep-ensure:
	@echo "==> Update go module dependencies..."
	export GO111MODULE=on; \
	go mod tidy

run-reader:
	@echo "==> Running local reader command..."
	export SCOPE=reader; \
	cd cmd/app/ && go run main.go

run-writer:
	@echo "==> Running local writer command..."
	export SCOPE=writer; \
	cd cmd/app/ && go run main.go

run-worker:
	@echo "==> Running local worker command..."
	export SCOPE=worker; \
	cd cmd/app/ && go run main.go

start-db:
	@echo "==> Starting database..."
	cd config && docker-compose up

start-db-async:
	@echo "==> Starting database..."
	cd config && docker-compose up -d

stop-db:
	@echo "==> Stopping database..."
	cd config && docker-compose down

check_fmt_linter_coverage: fmt linter test-cover

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: check_tools test test-cover fmt linter dep-ensure run-reader start-db start-db-async stop-db check_fmt_linter_coverage
