APP?=allure-go
RELEASE?=0.6.0
GOOS?=darwin

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

export GO111MODULE=on
export GOSUMDB=off
LOCAL_BIN:=$(CURDIR)/bin
EXAMPLES_TAGS:=examples_new,provider_new,allure_go_new,async

##################### GOLANG-CI RELATED CHECKS #####################
# Check global GOLANGCI-LINT
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.38.0

# Check local bin version
ifneq ($(wildcard $(GOLANGCI_BIN)),)
GOLANGCI_BIN_VERSION:=$(shell $(GOLANGCI_BIN) --version)
ifneq ($(GOLANGCI_BIN_VERSION),)
GOLANGCI_BIN_VERSION_SHORT:=$(shell echo "$(GOLANGCI_BIN_VERSION)" | sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_BIN_VERSION_SHORT:=0
endif
ifneq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_BIN_VERSION_SHORT)))"
GOLANGCI_BIN:=
endif
endif

# Check global bin version
ifneq (, $(shell which golangci-lint))
GOLANGCI_VERSION:=$(shell golangci-lint --version 2> /dev/null )
ifneq ($(GOLANGCI_VERSION),)
GOLANGCI_VERSION_SHORT:=$(shell echo "$(GOLANGCI_VERSION)"|sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_VERSION_SHORT:=0
endif
ifeq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_VERSION_SHORT)))"
GOLANGCI_BIN:=$(shell which golangci-lint)
endif
endif
##################### GOLANG-CI RELATED CHECKS #####################

.PHONY: full-demo
full-demo: install demo

.PHONY: demo
demo: examples allure-serve

.PHONY: install
install: .install_deps .install_allure

.PHONY: .install_deps
.install_deps:
	go mod tidy && go mod download

.PHONY: .install_allure
.install_allure:
ifeq ($(OS),Windows_NT)
	$(info Run Windows run pattern...)
	$(info Make sure scoop installed at your system. Check for more information: https://github.com/ScoopInstaller/Scoop#installation)
	scoop install allure
endif
ifeq ($(OS),Linux)
	$(info Run Linux run pattern...)
	$(info Make sure you have sudo rights for the system.)
	sudo apt-add-repository ppa:qameta/allure
	sudo apt-get update
	sudo apt-get install allure
endif
ifeq ($(OS),Darwin)
	$(info Run installation for Darwin OS)
	$(info Make sure brew installed at your system. Check for more information: https://docs.brew.sh/Installation)
	brew install allure
endif

.PHONY: examples
examples:
ifeq ($(OS),Windows_NT)
	$(info Run windows pattern...)
	set ALLURE_OUTPUT_PATH=../&& go test ./examples/... --tags=$(EXAMPLES_TAGS)
else
	$(info Run default pattern...)
	export ALLURE_OUTPUT_PATH=../ && go test ./examples/... --tags=$(EXAMPLES_TAGS) || true
endif

.PHONY: allure-serve
allure-serve:
	allure serve ./examples/allure-results

# run full lint like in pipeline
.PHONY: lint
lint: install-lint
	$(GOLANGCI_BIN) run --config=.golangci.yaml ./... --build-tags=$(EXAMPLES_TAGS)

.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info #Downloading golangci-lint v$(GOLANGCI_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG) && \
		go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint && \
		rm -rf $$tmp
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif
