PROJECT="uci-engine-wrapper"

build:
	@GOOS=darwin GOARCH=amd64 go build -o ../../dist/${PROJECT}.darwin-amd64 main.go

run: build
	@../../dist/${PROJECT}.darwin-amd64

