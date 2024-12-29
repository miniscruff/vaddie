.DEFAULT: help

.PHONY: help
help:
	@grep -E '^[a-z-]+:.*#' Makefile | \
		sort | \
		while read -r l; do printf "\033[1;32m$$(echo $$l | \
		cut -d':' -f1)\033[00m:$$(echo $$l | cut -d'#' -f2-)\n"; \
	done

.PHONY: test
test: # Run unit test suite
	go test -race -coverprofile=c.out ./...
	go tool cover -html=c.out -o=coverage.html

.PHONY: lint
lint: # Run linters
	go mod tidy
	golangci-lint run --fix ./...

.PHONY: format
format: lint # alias for lint

.PHONY: examples
examples: # Run the examples
	go run ./_examples/user
