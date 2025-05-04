build:
	@go build -o bin/GIN_API_V2 cmd/main.go

test:
	@go test ./...

vtest:
	@go test -v ./...


run: build
	@./bin/GIN_API_V2

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out `@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down