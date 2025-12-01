build:
	go build -o bin/gendiff ./cmd/gendiff
lint:
	golangci-lint run ./...
test:
	go test ./...
coverage:
	go test ./... -coverprofile=coverage.out