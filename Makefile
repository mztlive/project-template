# Makefile

# Get the current OS
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Set the output directories
HTTP_OUT_DIR := ./bin/$(OS)/http
CRONTAB_OUT_DIR := ./bin/$(OS)/crontab
RPC_OUT_DIR := ./bin/$(OS)/rpc
BACK_OUT_DIR := ./bin/$(OS)/back

# Set the output file paths
HTTP_APP := $(HTTP_OUT_DIR)/app
CRONTAB_APP := $(CRONTAB_OUT_DIR)/app
RPC_APP := $(RPC_OUT_DIR)/app
BACK_APP := $(BACK_OUT_DIR)/app

# current directory
CUR_DIR := $(shell pwd)

.PHONY: all build-http build-crontab copy-config run-http run-crontab clean

build: build-http build-crontab build-rpc build-back copy-config

build-http:
	@echo "Building http app for $(OS)..."
	@mkdir -p $(HTTP_OUT_DIR)
	@go build -o $(HTTP_APP) ./cmd/http/main.go

build-crontab:
	@echo "Building crontab app for $(OS)..."
	@mkdir -p $(CRONTAB_OUT_DIR)
	@go build -o $(CRONTAB_APP) ./cmd/crontab/main.go

build-rpc:
	@echo "Building rpc app for $(OS)..."
	@mkdir -p $(RPC_OUT_DIR)
	@go build -o $(RPC_APP) ./cmd/rpc/main.go

build-back:
	@echo "Building back app for $(OS)..."
	@mkdir -p $(BACK_OUT_DIR)
	@go build -o $(BACK_APP) ./cmd/back/main.go

copy-config:
	@echo "Copying local.toml to output directories..."
	@cp local.toml $(HTTP_OUT_DIR)
	@cp local.toml $(CRONTAB_OUT_DIR)
	@cp local.toml $(RPC_OUT_DIR)
	@cp local.toml $(BACK_OUT_DIR)

run-http:
	@echo "Running http app..."
	go run ./cmd/http/main.go -c $(CUR_DIR) -log ./log

run-crontab:
	@echo "Running http app..."
	go run ./cmd/crontab/main.go -c $(CUR_DIR) -log ./log

clean:
	@echo "Cleaning up build files..."
	@rm -rf ./bin