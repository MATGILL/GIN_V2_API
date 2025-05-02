build:
	@go build -o bin/GIN_API_V2 cmd/main.go

test:
	@go test §v ./...

run: build
	@./bin/GIN_API_V2