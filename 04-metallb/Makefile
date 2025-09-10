BUF		  := buf
GO		  := go
KO        := ko

CLUSTER   := dev-cluster
GEN_DIR	  := gen
BUILD_DIR := build
SERVICES  := booking
IMAGES    := $(SERVICES:%=%-img)	

SVC_BOOKING_IPv4 := $(shell kubectl get svc booking-service -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

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
		echo 'curl --show-error -s -H "Content-Type: application/json" \' ;\
		echo '     --data '\''{"passenger": "John Doe", "departure": "Paris", "destination": "Amsterdam"}'\'' \' ;\
		echo '     http://$${1:-127.0.0.1}:50051/booking.v1.BookingService/CreateBooking' ;\
	} > $(BUILD_DIR)/$@/booking-http-client
	@ echo "[INFO] Booking service listening on $(SVC_BOOKING_IPv4)"
	@chmod a+x $(BUILD_DIR)/$@/booking-http-client

clean: ## Clean build targets.
	@rm -rf $(GEN_DIR) $(BUILD_DIR)

images: $(IMAGES) ## Build service images.

$(IMAGES): 
	$(KO) build --local ./cmd/$(@:-img=)

deploy: ## Build and deploy service images to KinD.
	KO_DOCKER_REPO=kind.local KIND_CLUSTER_NAME=$(CLUSTER) $(KO) apply -f config/

delete: ## Teardown resources from KinD.
	$(KO) delete -f config/
