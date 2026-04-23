.PHONY: build build-windows run fmt test tidy clean

BINARY ?= cli2docker
WINDOWS_ARCH ?= amd64
WINDOWS_BINARY ?= $(BINARY)-windows-$(WINDOWS_ARCH).exe

build:
	go build -o $(BINARY) .

build-windows:
	GOOS=windows GOARCH=$(WINDOWS_ARCH) go build -o $(WINDOWS_BINARY) .

run:
	go run .

fmt:
	gofmt -w main.go

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -f $(BINARY) cli2docker-*.exe
