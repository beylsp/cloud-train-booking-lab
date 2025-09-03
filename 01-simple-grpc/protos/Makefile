.PHONY: lint
lint:
	@buf lint

.PHONY: format
format:
	@buf format -d -w

.PHONY: generate
generate: lint format
	@buf generate --clean
