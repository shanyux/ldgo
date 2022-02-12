#
# Copyright (C) distroy
#

# variables
PROJECT_ROOT= $(patsubst %/,%,$(abspath $(dir $$PWD)))
$(info PROJECT_ROOT: $(PROJECT_ROOT))

# go
# GO=env GODEBUG=madvdontneed=1 go
GO=go
GO_FLAGS=${flags}
GO_VERSION=$(shell go version | cut -d" " -f 3)
GO_MAJOR_VERSION=$(shell echo $(GO_VERSION) | cut -d"." -f 1)
GO_SUB_VERSION=$(shell echo $(GO_VERSION) | cut -d"." -f 2)
export GO111MODULE=on
# ifeq ($(shell expr ${GO_SUB_VERSION} '>' 10), 1)
# 	GO_FLAGS+=-mod=vendor
# endif
$(info GO_VERSION: $(GO_MAJOR_VERSION).$(GO_SUB_VERSION))
$(info GO_FLAGS: $(GO_FLAGS))

# go test
GO_TEST_DIRS+=$(shell find . -name '*_test.go' | grep -v -E 'vendor|bak' | xargs dirname | sort | uniq)
GO_TEST_DIRS_NAME=$(notdir $(GO_TEST_DIRS))
$(info GO_TEST_DIRS: $(GO_TEST_DIRS_NAME))

ifeq (${test_report},)
	export test_report=$(PROJECT_ROOT)/log
endif
# GO_TEST_FLAGS+=-v
GO_TEST_FLAGS+=-gcflags="-l -N"
GO_TEST_OUTPUT=${test_report}

# git
GIT_REVISION=$(shell git rev-parse HEAD 2> /dev/null)
GIT_BRANCH=$(shell git symbolic-ref HEAD 2> /dev/null | sed -e 's/refs\/heads\///')
GIT_TAG=$(shell git describe --exact-match --tags 2> /dev/null)
$(info GIT_REVISION: $(GIT_REVISION))
$(info GIT_BRANCH: $(GIT_BRANCH))
$(info GIT_TAG: $(GIT_TAG))

.PHONY: all
all: setup go-test

.PHONY: $(GO_TEST_DIRS_NAME)
$(GO_TEST_DIRS_NAME):
	@echo GO_TEST_DIRS: $(notdir $@)
	$(GO) test $(GO_FLAGS) $(GO_TEST_FLAGS) -v ./$(notdir $@)

.PHONY: dep
dep: setup
	$(GO) mod tidy
	# $(GO) mod vendor

.PHONY: go-test-coverage
go-test-coverage: setup
	@echo GO_TEST_DIRS: $(GO_TEST_DIRS_NAME)
	$(GO) test $(GO_FLAGS) $(GO_TEST_FLAGS) $(GO_TEST_DIRS) -json > "$(GO_TEST_OUTPUT)/test.json"
	$(GO) test $(GO_FLAGS) $(GO_TEST_FLAGS) $(GO_TEST_DIRS) -coverprofile="$(GO_TEST_OUTPUT)/coverage.out"

.PHONY: go-test
go-test: setup
	@echo GO_TEST_DIRS: $(GO_TEST_DIRS_NAME)
	$(GO) test $(GO_FLAGS) $(GO_TEST_FLAGS) -v $(GO_TEST_DIRS)

.PHONY: setup
setup:
	go install github.com/distroy/gocognit/cmd/gocognit@v1.0.5.2
	git config core.hooksPath "$(PROJECT_ROOT)/script/git-hooks"

.PHONY: complexity
complexity: setup
	gocognit -over 15 . >&2
