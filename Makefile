GO := go
FMT := gofmt
SHELL := /bin/bash
TARGET := $(shell echo $${PWD\#\#*/})
VERSION := 0.1
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
LDFLAGS = -ldflags "-X main.version=${VERSION}"

.PHONY: all
all: clean fmt build

fmt:
	$(FMT) -l -w $(SRC)

build: $(SRC)
	$(GO) build $(LDFLAGS) -o ./bin/$(TARGET) .

clean:
	rm -f $(TARGET) ./bin/$(TARGET)

vendor:
	govendor add +external
