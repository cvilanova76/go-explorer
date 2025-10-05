# Makefile for go-explorer

# Output binary directory
BIN_DIR := bin
BIN_NAME := go-explorer

# Command package path
CMD := ./cmd/cli

.PHONY: run build install clean

# Run the command with the COUNTRY variable (default: Portugal)
run:
	@COUNTRY=$(or $(COUNTRY),Portugal) \
	&& echo "Running with COUNTRY=$${COUNTRY}" \
	&& go run $(CMD) $${COUNTRY}

# Build into a binary in bin/
build:
	@mkdir -p $(BIN_DIR)
	@echo "Building $(BIN_NAME) into $(BIN_DIR)/$(BIN_NAME)"
	@go build -o $(BIN_DIR)/$(BIN_NAME) $(CMD)

# Install via 'go install'
install:
	@echo "Installing $(CMD)"
	@go install $(CMD)

clean:
	@echo "Cleaning $(BIN_DIR)"
	@rm -rf $(BIN_DIR)
