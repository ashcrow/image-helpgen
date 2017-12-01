VERSION := $(shell cat ./VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +%s)

# Used during all builds
LDFLAGS := -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildTime=${BUILD_TIME}


CONFIG_DIR ?= /etc
BIN_DIR ?= /usr/bin

.PHONY: help build clean deps install

help:
	@echo "Targets:"
	@echo " - build: Build the image-helpgen binary"
	@echo " - clean: Clean up after build"
	@echo " - deps: Install required tool and dependencies for building"
	@echo " - install: Install build results to the system"

build:
	go build -ldflags '${LDFLAGS}' -o image-helpgen cmd/main.go
	strip image-helpgen

clean:
	rm -f image-helpgen

deps:
	go get -u github.com/kardianos/govendor
	govendor sync

install:
	install -d ${PREFIX}${CONFIG_DIR}/image-helpgen/
	install --mode 644 template.tpl ${PREFIX}${CONFIG_DIR}/image-helpgen/template.tpl
	install --mode 755 image-helpgen ${PREFIX}${BIN_DIR}/image-helpgen

