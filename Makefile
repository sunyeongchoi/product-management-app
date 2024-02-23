$(GOPATH)/bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b `go env GOPATH`/bin v1.52.2

.PHONY: lint
lint: $(GOPATH)/bin/golangci-lint
	# Tidy Go modules
	go mod tidy
	# Lint Go files
	$(GOPATH)/bin/golangci-lint run --fix --verbose
