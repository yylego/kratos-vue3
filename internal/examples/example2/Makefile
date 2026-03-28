# Generate TypeScript gRPC client from proto files

API_PROTO_FILES = $(shell find ./api -name "*.proto")
OUTPUT_DIR = ./gen/ts

.PHONY: install
install:
	@echo "Installing protobuf-ts plugin..."
	npm install -g @protobuf-ts/plugin
	@echo "Installation complete!"

.PHONY: generate
generate:
	@echo "Generating TypeScript gRPC client code..."
	@mkdir -p $(OUTPUT_DIR)
	NODE_OPTIONS="--no-webstorage" \
	protoc \
		--plugin=protoc-gen-ts=$$(which protoc-gen-ts) \
		--ts_opt=ts_nocheck \
		--ts_opt=eslint_disable \
		--ts_opt=long_type_string \
		--ts_out=$(OUTPUT_DIR) \
		--proto_path=./api \
		--proto_path=../../proto3ps \
		$(API_PROTO_FILES)
	@echo "Generation complete! Output: $(OUTPUT_DIR)"

.PHONY: refresh
refresh:
	@echo "Updating test data from generated files..."
	cp $(OUTPUT_DIR)/v1/echo/echo.client.ts ./testdata/echo.client.ts-source.txt
	@echo "Updated source: testdata/echo.client.ts-source.txt"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Clean complete!"

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  install  - Install protobuf-ts plugin"
	@echo "  generate - Generate TypeScript gRPC client code"
	@echo "  refresh  - Update source test data"
	@echo "  clean    - Remove generated files"
	@echo "  help     - Show this help message"
