PROJECT_NAME := "opsgenie-exporter"
TARGET:=cmd/opsgenie-exporter/main.go
APP_INTERNAL_DEPTH := $(shell find ./internal -type d -printf '%d\n' | sort -rn | head -1)

.PHONY: all dep build clean lint

help: ## Show commands and description
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Install all dependencies
	@echo "Get the dependencies..."
	@make dep --silent 
	@echo "Install staticcheck to lint..."
	@go install honnef.co/go/tools/cmd/staticcheck@2022.1.2
#	@go isntall golang.org/x/lint/golint
# 	https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
	@echo "Done."


all: build

run-docker: dep-dev-run ## Run web application and dependencies inside container
	@docker-compose -f docker-compose.yml -f docker-compose.yml up -d  --build --remove-orphans

run: dep-dev-run ## Run web application and dependencies
	@ENVIRONMENT=development APP_VERSION="v-$(shell git rev-parse --short HEAD)" go run ${TARGET}

dep-dev-run: ## Run development dependencies
	@docker-compose up -d --build  --remove-orphans

dep-dev-stop: ## Stop development dependencies
	@docker-compose stop

dep-dev-status: ## Show status from development dependencies
	@docker-compose status

dev-start: dep-dev-run ## Run application and dependencies
	@docker-compose up --build  --remove-orphans

dev-destroy: dep-dev-stop ## Run application and dependencies
	@docker-compose down -v

lint: ## Lint the files
	@staticcheck ./... # TODO: Dosn't work with go 1.18
#	@golint ./... 

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -v -o bin/${PROJECT_NAME} ${TARGET}

clean: ## Remove previous build
	@rm -f bin/$(PROJECT_NAME)

scan-code:
	docker run --rm \
		-e HORUSEC_CLI_FILES_OR_PATHS_TO_IGNORE="*tmp*, **/.vscode/**, **/docs/**, **/node_modules/**, **/.horusec/**, **/.trivy/**" \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(PWD):/src horuszup/horusec-cli:latest \
		horusec start -p /src -P $(PWD)

scan-image:
	docker run --rm \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v ${PWD}/.trivy/.cache:/root/.cache/ \
		aquasec/trivy:0.18.3 \
		incident-webhook_incident-webhook

scan: scan-code scan-image
