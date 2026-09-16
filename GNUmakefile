TIMEOUT ?= 30s

default: test

fmt: ## Run gofmt
	gofmt -s -w ./

golangci-lint: ## Run golangci-lint
	@golangci-lint run ./...

test: ## Run unit tests
	go test -timeout=$(TIMEOUT) -parallel=4 ./...

tools: ## Install tools
	cd tools && go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint

# Please keep targets in alphabetical order
.PHONY: \
	fmt \
	golangci-lint \
	test \
	tools
