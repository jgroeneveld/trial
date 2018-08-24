default: test

test:
	go test -v ./assert

nice:
	golint ./... && go vet ./... && gofmt -s -w .
