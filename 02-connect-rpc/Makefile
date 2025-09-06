BUF		  := buf
GO		  := go

GEN_DIR	  := gen
BUILD_DIR := build
SERVICES  := booking

all: proto lint $(SERVICES) demo ## Build project (proto, services and demo).

.PHONY: help
help:  ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Lint Go code.
	$(GO) vet ./...
	$(GO) fmt ./...

.PHONY: proto
proto: ## Lint, format and build proto.
	$(BUF) lint
	$(BUF) format -w
	$(BUF) generate

.PHONY: services
services: $(SERVICES) ## Build services.

$(SERVICES):
	$(GO) build -o $(BUILD_DIR)/services/$@ ./cmd/$@

.PHONY: demo
demo: ## Build demo.
	$(GO) build -o $(BUILD_DIR)/$@/booking-grpc-client ./examples/booking-client
	@echo "build booking-http-client"
	@{ \
		echo '#!/bin/bash' ;\
		echo 'curl -s -H "Content-Type: application/json" \' ;\
		echo '     --data '\''{"passenger": "John Doe", "departure": "Paris", "destination": "Amsterdam"}'\'' \' ;\
		echo '     http://localhost:50051/booking.v1.BookingService/CreateBooking' ;\
	} > $(BUILD_DIR)/$@/booking-http-client
	@chmod a+x $(BUILD_DIR)/$@/booking-http-client

clean: ## Clean build targets.
	@rm -rf $(GEN_DIR) $(BUILD_DIR)
