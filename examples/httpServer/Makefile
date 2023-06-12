ifeq ($(PREFIX),)
	PREFIX := /usr/local
endif

.PHONY: build
build:
	go build -o httpServer .

.PHONY: clean
clean:
	rm -f httpServer

.PHONY: dev
dev:
	go run main.go

PHONY: test
test:
	go test ./...
