docs:
	godoc -http=:6060

lint:
	golangci-lint run -v ./...

test:
	go test -v ./... -cover
