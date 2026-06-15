.PHONY: run swagger dev-run test

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go

dev-run: swagger run

COVERAGE_EXCLUDE=mocks|main.go|test
COVERAGE_THRESHOLD = 50

test:
	@# Generate a comma-separated list of packages to cover, excluding cmd, docs, and mocks
	$(eval COVER_PKGS := $(shell go list ./... | grep -vE "cmd|docs|mocks" | tr '\n' ',' | sed 's/,$$//'))
#
	@# Run tests only on actual code packages, using the filtered list for -coverpkg
	go test ./... -coverprofile=coverage.tmp -covermode=atomic -coverpkg=$(COVER_PKGS) -p 1
	#go test ./... -coverprofile=coverage.tmp -covermode=atomic -coverpkg=./... -p 1
	grep -vE "$(COVERAGE_EXCLUDE)" coverage.tmp > coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@total=$$(go tool cover -func=coverage.out | grep total: | awk '{print $$3}' | tr -d '%'); \
	if [ $$(echo "$$total < $(COVERAGE_THRESHOLD)" | bc -l) -eq 1 ]; then \
		echo "Coverage ($$total%) is below threshold ($(COVERAGE_THRESHOLD)%)"; \
		exit 1; \
	else \
		echo "Coverage ($$total%) meets threshold ($(COVERAGE_THRESHOLD)%)"; \
	fi
