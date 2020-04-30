# go commands
GOCMD := go
GOTEST := $(GOCMD) test
GOTOOL := $(GOCMD) tool

COVERPROFILE ?= /tmp/profile.out

.PHONY: test
test:
	$(GOTEST) ./...

.PHONY: coverage
coverage:
	$(GOTEST) -covermode=atomic -coverprofile=$(COVERPROFILE) ./...

.PHONY: showcoverage
showcoverage: coverage
	$(GOTOOL) cover -html=$(COVERPROFILE)
