.PHONY: run-server
run-server:
	@go mod tidy
	@go run cmd/server/*.go

.PHONY: build-server
build-server:
	@go mod tidy
	@go build -o bin/server cmd/server/*.go

.PHONY: run-client
run-client:
	@go mod tidy
	@go run cmd/client/*.go

.PHONY: build-client
build-client:
	@go mod tidy
	@go build -o bin/client cmd/client/*.go