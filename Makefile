GO        ?= GOSUMDB=off go
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin


.PHONY: build
build:
	GOBIN=$(BINDIR) $(GO) build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/davidbolet/gopasc/...

.PHONY: test
test: build
test: check-style
test: test-unit

prereqs:
	$(GO) mod tidy
	$(GO) mod download
	# Make sure there is no go.sum leftover as this would break the integration tests in docker
	rm go.sum

.PHONY: check-style
check-style: prereqs
	$(GO) fmt ./...
	$(GO) vet ./...

.PHONY: test-unit
test-unit: prereqs
	@echo
	@echo "==> Running unit tests <=="
	$(GO) test $(TESTFLAGS) `go list ./... | grep -v integration`

.PHONY: test-unit-cover
test-unit-cover: prereqs
	@echo
	@echo "==> Running unit tests with coverage<=="
	@echo "==> Check your browser for more details<=="
	@scripts/coverage.sh


.PHONY: gosec
gosec:
	GOSUMDB=off gosec ./...	

.PHONY: gosec-install	
gosec-install:
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin
