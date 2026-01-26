.PHONY: build run fmt test tidy clean

build:
	go build -o node2docker-go .

run:
	go run .

fmt:
	gofmt -w main.go

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -f node2docker-go
