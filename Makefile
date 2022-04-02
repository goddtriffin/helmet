$(VERBOSE).SILENT:
.DEFAULT_GOAL := help

COVERPROFILE ?= /tmp/profile.out

.PHONY: help
help: # displays Makefile target info
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
			IFS=$$':' ; \
			help_split=($$help_line) ; \
			help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			printf '\033[36m'; \
			printf "%-30s %s" $$help_command ; \
			printf '\033[0m'; \
			printf "%s\n" $$help_info; \
	done

.PHONY: test
test: ## runs all tests
	go test ./...

.PHONY: coverage
coverage: ## runs all tests with coverage
	go test -covermode=atomic -coverprofile=$(COVERPROFILE) ./...

.PHONY: showcoverage
showcoverage: coverage ## runs tests with coverage, then opens coverage file in browser
	go tool cover -html=$(COVERPROFILE)
