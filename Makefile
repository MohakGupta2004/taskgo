BINARY_NAME=taskgo
BUILD_DIR=bin

.PHONY: all build clean install run test

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

install: build
	@echo "Installing to /usr/local/bin..."
	@sudo mv $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete."

run: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

test:
	@go test ./...
