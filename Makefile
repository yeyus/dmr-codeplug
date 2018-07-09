TEST_RESULTS = ./test_results

.PHONY: help proto

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run the test suite for this package
	@mkdir -p $(TEST_RESULTS) && touch $(TEST_RESULTS)/coverage.out
	go test ./... -coverprofile=$(TEST_RESULTS)/coverage.out
	@go tool cover -html=$(TEST_RESULTS)/coverage.out -o $(TEST_RESULTS)/coverage.html

proto: ## Refresh go generated protos
	protoc --go_out=import_path=proto/tytera:. proto/tytera/*.proto
