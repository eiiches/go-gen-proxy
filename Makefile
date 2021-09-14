.PHONY: all
all: test build

.PHONY: test
test:
	$(RM) ./internal/test/generated.go
	go generate ./internal/test
	go test ./internal/test

.PHONY: build
build:
	go build ./cmd/go-gen-proxy

.PHONY: clean
clean:
	$(RM) ./go-gen-proxy
	$(RM) ./internal/test/generated.go
