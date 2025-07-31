test:
	@go test ./...

run: test
	@go run main.go