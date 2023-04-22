BINARY_NAME=todoInGo

coverage:
	go test -c ./ -cover -covermode=set -coverpkg=./... -o bin/$(BINARY_NAME) -tags=integration
	./bin/$(BINARY_NAME).test -test.coverprofile coverage.integration.out

# executes unit tests and captures coverage per function
test: coverage.profile

# executes unit test and captures coverage
coverage.out:
	go test ./... -coverprofile=coverage.out

build: bin/$(BINARY_NAME) bin/$(BINARY_NAME).test

bin/$(BINARY_NAME):
	go build -o bin/$(BINARY_NAME) .

bin/$(BINARY_NAME).test:
	go test -c ./ -cover -covermode=set -coverpkg=./... -o bin/$(BINARY_NAME).test -tags=integration

# computes unit test coverage per function
coverage.profile: coverage.out
	go tool cover -func=coverage.out -o coverage.profile

run-test: bin/$(BINARY_NAME).test
	./bin/$(BINARY_NAME).test -test.coverprofile coverage.out.integration SEPARATOR --output-mode overwrite

server:
	go run main.go

.PHONY: build run-test test server coverage
