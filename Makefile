SHELL := /bin/bash

APP_NAME := groupie-tracker
BIN_DIR := bin
GO := go
PORT ?= 8080
APIPORT ?= 8081
ENV_FILE ?= .env

RUN_ENV := PORT=$(PORT) APIPORT=$(APIPORT)

define run_with_env
@set -a; \
	[ -f $(ENV_FILE) ] && . $(ENV_FILE); \
	set +a; \
	$(1)
endef

.PHONY: help run build test fmt tidy clean

.DEFAULT_GOAL := help

help:
	@printf "Usage:\n"
	@printf "  make run       # Run the web + API servers with go run\n"
	@printf "  make build     # Build the binary into $(BIN_DIR)/$(APP_NAME)\n"
	@printf "  make test      # Run go test ./...\n"
	@printf "  make fmt       # Format Go files\n"
	@printf "  make tidy      # Update go.mod/go.sum\n"
	@printf "  make clean     # Remove build artifacts\n"

run:
	$(call run_with_env,$(RUN_ENV) $(GO) run .)

build:
	@mkdir -p $(BIN_DIR)
	$(call run_with_env,$(RUN_ENV) $(GO) build -o $(BIN_DIR)/$(APP_NAME) .)

test:
	$(call run_with_env,$(GO) test ./...)

fmt:
	$(call run_with_env,$(GO) fmt ./...)

tidy:
	$(call run_with_env,$(GO) mod tidy)

clean:
	rm -rf $(BIN_DIR)
