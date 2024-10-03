BINARY_NAME=anansi-profiler

clean:
	rm -f output.md && rm -f output.txt && rm -f bin/* && rm -f ${BINARY_NAME} && rm -f main

build:
	go get ./... && go build -o ${BINARY_NAME} cmd/main.go

build-bins:
	GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 cmd/main.go

start:
	go run cmd/main.go -n 20 -c 20 -r text --log-level="debug" sample.log

test:
	go test ./...

test_cover:
	go test -cover ./...

test_cover_html:
	go test -coverprofile=resources/coverage.out ./... && go tool cover -html=resources/coverage.out