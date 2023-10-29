ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.EXPORT_ALL_VARIABLES:
.PHONY: help
.DEFAULT_GOAL := help

help: ## 💬 This help message :)
	@figlet $@ || true
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## 🚀 Run the server
	@figlet $@ || true
	air

build: ## 🔨 Build the server
	@figlet $@ || true
	go build -o ./bin/server htmx-go-todo/server