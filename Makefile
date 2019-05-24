PROJECT_NAME := HOTEL AUTOMATION

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build application binaries
	GOOS=darwin GOARCH=amd64 go build -o bin/hotelautomation-darwin-amd64 .
	GOOS=linux GOARCH=amd64 go build -o bin/hotelautomation-linux-amd64 .

.PHONY: lint
lint: ## check code for lint errors
	go vet ./...

.PHONY: test
test: ## run unit tests
	go test -race ./...
