BUF		  := buf
GO		  := go

GEN_DIR	  := gen
BUILD_DIR := build
SERVICES  := booking api

all: lint proto $(SERVICES)

.PHONY: lint
lint:
	$(BUF) lint
	$(BUF) format -w
	$(GO) vet ./...
	$(GO) fmt ./...

.PHONY: proto
proto:
	$(BUF) generate

.PHONY: services
services: $(SERVICES)

$(SERVICES):
	$(GO) build -o $(BUILD_DIR)/$@ ./cmd/$@

clean:
	@rm -rf $(GEN_DIR) $(BUILD_DIR)
