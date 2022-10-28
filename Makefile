run:
	go run cmd/main.go
test:
	GIN_MODE=release go test -v ./...