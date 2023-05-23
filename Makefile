# Makefile

# Get the current OS
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Set the output directories
HTTP_OUT_DIR := ./bin/$(OS)/http
CRONTAB_OUT_DIR := ./bin/$(OS)/crontab

# Set the output file paths
HTTP_APP := $(HTTP_OUT_DIR)/app
CRONTAB_APP := $(CRONTAB_OUT_DIR)/app

.PHONY: all build-http build-crontab copy-config clean

build: build-http build-crontab copy-config

build-http:
	@echo "Building http app for $(OS)..."
	@mkdir -p $(HTTP_OUT_DIR)
	@go build -o $(HTTP_APP) ./cmd/http/main.go

build-crontab:
	@echo "Building crontab app for $(OS)..."
	@mkdir -p $(CRONTAB_OUT_DIR)
	@go build -o $(CRONTAB_APP) ./cmd/crontab/main.go

copy-config:
	@echo "Copying local.toml to output directories..."
	@cp local.toml $(HTTP_OUT_DIR)
	@cp local.toml $(CRONTAB_OUT_DIR)

clean:
	@echo "Cleaning up build files..."
	@rm -rf ./bin