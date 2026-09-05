.PHONY: test fmt tidy help

GO ?= go

test:
	$(GO) test -race -count=1 -timeout 60s ./...

fmt:
	$(GO) fmt ./...

tidy:
	$(GO) mod tidy

help:
	@echo "Targets:"
	@echo "  test  - run tests with race detection"
	@echo "  fmt   - gofmt all packages"
	@echo "  tidy  - go mod tidy"
