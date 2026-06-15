.PHONY: run swagger dev-run test

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go

dev-run: swagger run

COVERAGE_EXCLUDE=mocks|main.go|test|pkg
COVERAGE_THRESHOLD = 40

test:
	$(eval COVER_PKGS := $(shell go list ./... | grep -vE "cmd|docs|mocks" | tr '\n' ',' | sed 's/,$$//'))
	$(eval TEST_PKGS  := $(shell go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' ./... | grep -vE "cmd|docs|mocks"))
	go test $(TEST_PKGS) -coverprofile=coverage.tmp -covermode=atomic -coverpkg=$(COVER_PKGS) -p 1
	grep -vE "$(COVERAGE_EXCLUDE)" coverage.tmp > coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@total=$$(go tool cover -func=coverage.out | grep total: | awk '{print $$3}' | tr -d '%'); \
	if [ $$(echo "$$total < $(COVERAGE_THRESHOLD)" | bc -l) -eq 1 ]; then \
	   echo "Coverage ($$total%) is below threshold ($(COVERAGE_THRESHOLD)%)"; \
	   exit 1; \
	else \
	   echo "Coverage ($$total%) meets threshold ($(COVERAGE_THRESHOLD)%)"; \
	fi
