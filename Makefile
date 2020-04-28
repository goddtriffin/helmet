# go commands
GOCMD := go
GOTEST := $(GOCMD) test
GOTOOL := $(GOCMD) tool

COVERPROFILE ?= /tmp/profile.out

.PHONY: test
test:
	$(GOTEST) ./...

coverage:
	$(GOTEST) -covermode=atomic -coverprofile=$(COVERPROFILE) ./...

showcoverage: coverage
	$(GOTOOL) cover -html=$(COVERPROFILE)
