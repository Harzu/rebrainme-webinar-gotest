.PHONY: generate mock-gen build run-test-env migration-up migration-down run-long-tests stop-env

export APP_NAME := gotest
export VERSION := $(if $(TAG),$(TAG),$(if $(BRANCH_NAME),$(BRANCH_NAME),$(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)))
export DOCKER_REPOSITORY := rebrainme-webinars
export DOCKER_BUILDKIT := 1

MIGRATE_DSN := "postgres://gotest:gotest@psql:5432/gotest?sslmode=disable"
NOCACHE := $(if $(NOCACHE),"--no-cache")

generate: mock-gen

mock-gen:
	@rm -rf ./test/mocks/packages
	@go generate ./...

build:
	@docker build ${NOCACHE} --pull -f ./build/helper.Dockerfile -t ${DOCKER_REPOSITORY}/${APP_NAME}-helper:${VERSION} --ssh default .

run-test-env:
	@docker-compose -f ./build/docker-compose.yml up -d psql
	@sleep 5

migration-up:
	@docker-compose -f ./build/docker-compose.yml run --rm -T helper migrate -verbose -path ./migrations -database ${MIGRATE_DSN} up

migration-down:
	@docker-compose -f ./build/docker-compose.yml run --rm -T helper migrate -verbose -path ./migrations -database ${MIGRATE_DSN} down

tests-short:
	@go test -short -cover ./...

tests-long: run-test-env migration-up
	@docker-compose -f ./build/docker-compose.yml run --rm -T helper sh ./scripts/long_tests_runner.sh; $(MAKE) stop-env

stop-env:
	@docker-compose -f ./build/docker-compose.yml down