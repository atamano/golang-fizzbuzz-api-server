  
PACKAGES := $(shell go list ./... | grep -v /vendor/)

APP_CONTAINER_NAME=server
DOCKER_COMPOSE := docker-compose -f docker/docker-compose.yml
EXEC ?=$(DOCKER_COMPOSE) exec -T $(APP_CONTAINER_NAME)

.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	${EXEC} go tool cover -html=coverage-all.out

.PHONY: migration-init
migration-init: ## add migration table
	${EXEC} go run ./cmd/migrate/*.go init

.PHONY: migration-up
migration-up: ## run mirations
	${EXEC} go run ./cmd/migrate/*.go up

.PHONY: migration-down
migration-down: ## down mirations
	${EXEC} go run ./cmd/migrate/*.go down

.PHONY: restart-server
restart-server: ## restart server
	$(DOCKER_COMPOSE) restart ${APP_CONTAINER_NAME}

.PHONY: start
start: ## up docker-compose services
	$(DOCKER_COMPOSE) up -d --build

.PHONY: livereload
livereload: ## livereload on change (needs fswatch installed)
	fswatch -x -o --event Created --event Updated --event Renamed -r internal pkg cmd | xargs -n1 -I {} make restart-server

.PHONY: stop
stop:  ## remove docker containers
	$(DOCKER_COMPOSE) kill
	$(DOCKER_COMPOSE) rm -v --force