.PHONY: test
test:
	go test -v ./...

.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.out ./...

.PHONY: test-cover-html
test-cover-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: remove-cover
remove-cover:
	grep -v -E '_stub|_mock' coverage.out > coverage_cleaned.out

.PHONY: cover
cover:
	go tool cover -func=coverage_cleaned.out

.PHONY: lint
lint:
	staticcheck ./...
	



