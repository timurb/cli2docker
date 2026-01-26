.PHONY: build run fmt test tidy clean

build:
	go build -o cli2docker .

run:
	go run .

fmt:
	gofmt -w main.go

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -f cli2docker
