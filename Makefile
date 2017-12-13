VERSION := $(shell cat ./VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +%s)
DEFAULT_TEMPLATE := /etc/image-helpgen/template.tpl

# Used during all builds
LDFLAGS := -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildTime=${BUILD_TIME} -X main.defaultTemplate=${PREFIX}${DEFAULT_TEMPLATE}

CONFIG_DIR ?= /etc
BIN_DIR ?= /usr/bin

.PHONY: help build clean deps install lint test

help:
	@echo "Targets:"
	@echo " - build: Build the image-helpgen binary"
	@echo " - clean: Clean up after build"
	@echo " - deps: Install required tool and dependencies for building"
	@echo " - install: Install build results to the system"
	@echo " - lint: Run golint"
	@echo " - test: Run unittests"
	@echo ""
	@echo "Variables:"
	@echo " - PREFIX: The root location to install. This prepends to all *_DIR variables. Set to: ${PREFIX}"
	@echo " - CONFIG_DIR: The directory that houses configuration files. Set to: ${CONFIG_DIR}"
	@echo " - BIN_DIR: The directory that houses binaries. Set to: ${BIN_DIR}"
	@echo " - DEFAULT_TEMPLATE: The default template to load if none is provided. Set to: ${DEFAULT_TEMPLATE}"
	@echo " - VERSION: Generally not overridden. The output of the VERSION file. Set to: ${VERSION}"
	@echo " - COMMIT_HASH: Generally not overridden. The git hash the code was built from. Set to: ${COMMIT_HASH}"
	@echo " - BUILD_TIME: Generally not overridden. The unix time of the build. Set to: ${BUILD_TIME}"

build:
	go build -ldflags '${LDFLAGS}' -o image-helpgen main.go
	strip image-helpgen

clean:
	rm -f image-helpgen

deps:
	go get -u github.com/kardianos/govendor
	govendor sync

install: clean build
	install -d ${PREFIX}${CONFIG_DIR}/image-helpgen/
	install --mode 644 template.tpl ${PREFIX}${CONFIG_DIR}/image-helpgen/template.tpl
	install -d ${PREFIX}${BIN_DIR}
	install --mode 755 image-helpgen ${PREFIX}${BIN_DIR}/image-helpgen

lint:
	go get -u github.com/golang/lint/golint
	golint . cmd/ types/ utils/


test:
	go list ./... | grep -v vendor | xargs govendor test -v

e2e: build
	./e2e.sh
