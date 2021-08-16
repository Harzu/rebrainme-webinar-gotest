SHELL := bash
.ONESHELL:
MAKEFLAGS += --no-builtin-rules

export APP_NAME := gotest
export VERSION := $(if $(TAG),$(TAG),$(if $(BRANCH_NAME),$(BRANCH_NAME),$(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)))
export DOCKER_REPOSITORY := rebrainme-webinars
export DOCKER_BUILDKIT := 1

MIGRATE_DSN := "postgres://gotest:gotest@postgres:5432/gotest?sslmode=disable"
NOCACHE := $(if $(NOCACHE),"--no-cache")

.PHONY: help
help: ## List all available targets with help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Init project
	@git config core.hooksPath .githooks
	@go mod tidy && make generate

.PHONY: generate
generate: ## Golang codegen
	@go generate ./...

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run

.PHONY: build
build: ## Build docker containers
	@docker build ${NOCACHE} --pull -f ./build/helper.Dockerfile -t ${DOCKER_REPOSITORY}/${APP_NAME}-helper:${VERSION} --ssh default .

.PHONY: run-dev-env
run-dev-env:
	@docker-compose up -d postgres

.PHONY: migration-up
migration-up: ## Run develop migrations
	@docker-compose run --rm -T helper migrate -verbose -path ./migrations -database ${MIGRATE_DSN} up

.PHONY: migration-down
migration-down: ## Rollback develop migrations
	@docker-compose run --rm helper migrate -verbose -path ./migrations -database ${MIGRATE_DSN} down

.PHONY: test-short
test-short: ## Run only unit tests
	@go test -short -cover ./...

.PHONY: test-long
test-long: run-dev-env ## Run all tests (unit/integrations)
	@make run-dev-env && make migration-up && make test-long-up; make stop

.PHONY: tests-long-up
test-long-up:
	@docker-compose run --rm helper sh ./scripts/long_tests_runner.sh; $(MAKE) stop-env

.PHONY: stop
stop: ## Stop dev environment
	@docker-compose down -v