docs:
	godoc -http=:6060

lint:
	golangci-lint run -v ./...

run:
	@go build -o cmd/$(cmd)/$(cmd) cmd/$(cmd)/main.go

test:
	go test -v ./... -cover
