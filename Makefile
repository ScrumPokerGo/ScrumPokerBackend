.PHONY: build build-alpine clean test help default

BIN_NAME=ScrumPoker

VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
GOPATH   = $(CURDIR)
PACKAGE  =  ScrumPokerBackend
BASE     = $(GOPATH)/src/$(PACKAGE)/

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build "${BASE}main.go"

GLIDE = glide

glide.lock: glide.yaml | $(BASE)
	cd $(BASE) && $(GLIDE) update
	@touch $@
vendor: glide.lock | $(BASE)
	cd $(BASE) && $(GLIDE) --quiet install
	@ln -sf . vendor/src
	@touch $@
