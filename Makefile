test:
	go test -v -cover -short ./...

coverage:
	go test -c ./ -cover -covermode=set -coverpkg=./... -tags=integration
	./todoInGo.test -test.coverprofile coverage.integration.out
server:
	go run main.go

.PHONY: test server coverage
