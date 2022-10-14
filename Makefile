AWS_PROFILE=default
AWS_REGION=us-east-1
STAGE=dev
PROJECT=${STAGE}-ctx-kitchensink-go

CSV_FILE=${CSV_FILE_1000_ITEM}

CSV_FILE_1_ITEM=../../extra/data/the_movie_database/movies_metadata_1.csv
CSV_FILE_50_ITEM=../../extra/data/the_movie_database/movies_metadata_50.csv
CSV_FILE_1000_ITEM=../../extra/data/the_movie_database/movies_metadata_1000.csv
CSV_FILE_ALL_ITEM=../../extra/data/the_movie_database/movies_metadata.csv	# NOTE: Fast import needs work for this much data.

RESOURCE_URL := ${shell grep MoviesApiEndpoint config/local-dev.json|cut -f4 -d\"}

DATA_IMPORT_CMD_PATH=cmd/cli/import_data/main.go

LINT_PATH=reports/lint
LINT_FILE=lint_report_$(shell date '+%Y-%m-%d-%H%M%S').html


ifeq ($(STAGE), local)
	DEPLOY_OS   := ${shell uname -s|tr '[:upper:]' '[:lower:]'}
	DEPLOY_ARCH := amd64
#	DEPLOY_ARCH := ${shell uname -m} # hard code to "amd64" for MacOS fix
else
	DEPLOY_OS   := linux
	DEPLOY_ARCH := amd64
endif

DEPLOY_OS   := darwin
DEPLOY_ARCH := amd64


COVERAGE_PATH=reports/coverage
COVERAGE_FILE=coverage_report_$(shell date '+%Y-%m-%d-%H%M%S').html

MODULE_PATH=backend/mainmodule

# https://image.tmdb.org/t/p/w185/iYM0EdHrcbu8rhkDKVBkDcZEo8t.jpg?api_key=5e0466ac7c5f3df99a54a43382b798e0
# https://image.tmdb.org/t/p/w185_and_h278_bestv2/iYM0EdHrcbu8rhkDKVBkDcZEo8t.jpg?api_key=5e0466ac7c5f3df99a54a43382b798e0
# HELP
.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo "\n\n------------------------------------------\Make Command Help\n\n"
	@echo "\nGet more detail on SST here: https://docs.serverless-stack.com/packages/cli\n"
	@echo "\nTARGETS:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

# TARGETS
install: 
	npm install
	@echo "install golang https://go.dev/doc/install"

build: config/local-${STAGE}.json	## Build project for deployment
	AWS_PROFILE=${AWS_PROFILE} GOOS=${DEPLOY_OS} GOARCH=${DEPLOY_ARCH} npx sst build --stage ${STAGE}

start:	## Start "local" environment with supporting cloud resources
	make start-local STAGE=local

start-${STAGE}: config/local-${STAGE}.json	## Start environment with supporting cloud resources
	AWS_PROFILE=${AWS_PROFILE} GOOS=${DEPLOY_OS} GOARCH=${DEPLOY_ARCH} npx sst start --stage ${STAGE}

deploy:	config/local-${STAGE}.json ## Deploy stack as complete build to AWS environment
	@echo "Deploy Resource"
	AWS_PROFILE=${AWS_PROFILE} GOOS=${DEPLOY_OS} GOARCH=${DEPLOY_ARCH} npx sst deploy --stage ${STAGE}

config/local-${STAGE}.json: config/deploy.json
	make update-config

update-config: ## Update AppConfig with latest stack details
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=${AWS_REGION} node scripts/update-config.js ${STAGE} ${PROJECT}

deploy-config: ## Deploy AppConfig with latest stack details
	AWS_PROFILE=${AWS_PROFILE} GOOS=${DEPLOY_OS} GOARCH=${DEPLOY_ARCH} npx sst deploy --stage ${STAGE} config-stack

remove:	## Remove environment
	AWS_PROFILE=${AWS_PROFILE} npx sst remove --stage ${STAGE}

remove-all:	## Remove ALL environments including local
	make remove-local STAGE=local
	make remove-${STAGE}
	rm -Rf .build

data-import:	## Import test data into environment
	@echo "WARNING: Update RESOURCE_URL variable if needed."
	cd backend/mainmodule; go run ${DATA_IMPORT_CMD_PATH} -table=${RESOURCE_URL} -path=${CSV_FILE}

data-import-fast:	## Fast Import test data into environment
	cd backend/mainmodule; go run ${DATA_IMPORT_CMD_PATH} -table=${RESOURCE_URL} -path=${CSV_FILE} -fast=true

mod-tidy:	## Tidy up the go module dependencies
	# Clean up any uncessessary go modules in go.mod
	go mod tidy

sst-console:	## View the SST Console
	AWS_PROFILE=${AWS_PROFILE} npx sst console --stage ${STAGE}

lint-install: ## install lining executable tool, golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint-run: ## Run lint without the install
	$(eval GP=$(shell go env GOPATH))
	$(eval XP=$(shell echo '../../node_modules/xunit-viewer/bin'))
	cd $(MODULE_PATH); $(GP)/bin/golangci-lint --out-format junit-xml run ./... > tmp.xml; $(XP)/xunit-viewer -r tmp.xml -o ../../$(LINT_PATH)/$(LINT_FILE); rm -rf tmp.xml
	cp -f $(LINT_PATH)/$(LINT_FILE) $(LINT_PATH)/lint_report_latest.html

lint: ## Run both lint install and run
	make lint-install && make lint-run

test-install:	## Install test packages
	cd $(MODULE_PATH); go get -u github.com/stretchr/testify
	go install github.com/axw/gocov/gocov@latest
	go install github.com/AlekSi/gocov-xml@latest
	
test-run: ## Run test without the install
	cd $(MODULE_PATH); go test -coverprofile=../../$(COVERAGE_PATH)/coverage.out -coverpkg=./... ./tests/... -v
	cd $(MODULE_PATH); go tool cover -html=../../$(COVERAGE_PATH)/coverage.out -o ../../$(COVERAGE_PATH)/$(COVERAGE_FILE)
	cd $(MODULE_PATH); gocov convert ../../$(COVERAGE_PATH)/coverage.out | gocov-xml > ../../$(COVERAGE_PATH)/coverage_report.xml
	cp -f $(COVERAGE_PATH)/$(COVERAGE_FILE) $(COVERAGE_PATH)/coverage_report_latest.html
	rm $(COVERAGE_PATH)/coverage.out
	@echo "\033[1;32mCoverage report available at $(COVERAGE_PATH)/$(COVERAGE_FILE)\033[0m"
	
test:	## Install and run tests
	make test-install && make test-run
	
godocs: ## Browse godocs for project using local http server
	@echo "NOTE: Docs will be hosted on http://127.0.0.1:6060"
	cd $(MODULE_PATH);godoc -http=127.0.0.1:6060;cd -
