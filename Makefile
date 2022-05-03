AWS_PROFILE=default
CSV_FILE=${CSV_FILE_1000_ITEM}

CSV_FILE_1_ITEM=../../extra/data/the_movie_database/movies_metadata_1.csv
CSV_FILE_50_ITEM=../../extra/data/the_movie_database/movies_metadata_50.csv
CSV_FILE_1000_ITEM=../../extra/data/the_movie_database/movies_metadata_1000.csv
CSV_FILE_ALL_ITEM=../../extra/data/the_movie_database/movies_metadata.csv	# NOTE: Fast import needs work for this much data.

RESOURCE_URL_LOCAL=https://u4vzxgzv0a.execute-api.us-east-1.amazonaws.com/movies
RESOURCE_URL_DEV=https://fpnl6cvn76.execute-api.us-east-1.amazonaws.com/movies

DATA_IMPORT_CMD_PATH=cmd/cli/import_data/main.go

LINT_PATH=reports/lint
LINT_FILE=lint_report_$(shell date '+%Y-%m-%d-%H%M%S').html

DEPLOY_OS=linux
DEPLOY_ARCH=amd64

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
start:	## Alias: Start "local" environment with supporting cloud resources
	make deploy-local
build:	## Build project for deployment
	AWS_PROFILE=${AWS_PROFILE} GOOS=${DEPLOY_OS} GOARCH=${DEPLOY_ARCH} npx sst build
deploy-local:	## Start "local" environment with supporting cloud resources, NOTE: this will use "local" stage prefix
	AWS_PROFILE=${AWS_PROFILE} npx sst start --stage local
deploy-dev:	## Deploy stack as complete build to AWS environment, NOTE: this will use "dev" stage prefix
	AWS_PROFILE=${AWS_PROFILE} GOOS=${DEPLOY_OS} GOARCH=${DEPLOY_ARCH} npx sst deploy --stage dev
deploy:	## Alias: Deploy stack as complete build to AWS environment, NOTE: this willuse "dev" stage prefix
	make deploy-dev
remove-local:	## Remove "local" environment and supporing cloud resources
	AWS_PROFILE=${AWS_PROFILE} npx sst remove --stage local
remove-dev:	## Remove "dev" environment
	AWS_PROFILE=${AWS_PROFILE} npx sst remove --stage dev
remove-all:	## Remove ALL environments including local
	make remove-local
	make remove-dev
	rm -Rf .build
data-import-local:	## Import test data into local environment
	@echo "WARNING: Update RESOURCE_URL variable if needed."
	cd backend/mainmodule; go run ${DATA_IMPORT_CMD_PATH} -u=${RESOURCE_URL_LOCAL} -p=${CSV_FILE}
data-import-dev:	## Import test data into production environment
	@echo "WARNING: Update RESOURCE_URL variable if needed."
	cd backend/mainmodule; go run ${DATA_IMPORT_CMD_PATH} -u=${RESOURCE_URL_DEV} -p=${CSV_FILE}
data-import-dev-fast:	## Fast Import test data into production environment
	cd backend/mainmodule; go run ${DATA_IMPORT_CMD_PATH} -f -u=${RESOURCE_URL_DEV} -p=${CSV_FILE}
mod-tidy:	## Tidy up the go module dependencies
	# Clean up any uncessessary go modules in go.mod
	go mod tidy
sst-console-local:	## View the SST Console
	AWS_PROFILE=${AWS_PROFILE} npx sst console --stage local
sst-console-dev:	## View the SST Console
	AWS_PROFILE=${AWS_PROFILE} npx sst console --stage dev
lint-install: ## install lining executable tool, golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint-run: ## Run lint without the install
	$(eval GP=$(shell go env GOPATH))
	$(eval XP=$(shell echo '../../node_modules/xunit-viewer/bin'))
	cd $(MODULE_PATH); $(GP)/bin/golangci-lint --out-format junit-xml run ./... > tmp.xml; $(XP)/xunit-viewer -r tmp.xml -o ../../$(LINT_PATH)/$(LINT_FILE); rm -rf tmp.xml
	cp -f $(LINT_PATH)/$(LINT_FILE) $(LINT_PATH)/lint_report_latest.html
lint: ## run both lint install and run
	make lint-install && make lint-run
test-install:	## install test packages
	#go get -u github.com/stretchr/testify
	go install github.com/jstemmer/go-junit-report@latest
test-run:	## run tests
	#cd $(MODULE_PATH); go test -coverprofile=../../$(COVERAGE_PATH)/coverage.out -coverpkg=./... ./tests/... -v
	#cd $(MODULE_PATH); go tool cover -html=../../$(COVERAGE_PATH)/coverage.out -o ../../$(COVERAGE_PATH)/$(COVERAGE_FILE)
	#cp -f $(COVERAGE_PATH)/$(COVERAGE_FILE) $(COVERAGE_PATH)/coverage_report_latest.html
	#rm $(COVERAGE_PATH)/coverage.out
	cd $(MODULE_PATH); go test -v 2>&1 | go-junit-report > ../../$(COVERAGE_PATH)/report.xml
	cat $(COVERAGE_PATH)/report.xml
	#@echo "\033[1;32mCoverage report available at $(COVERAGE_PATH)/$(COVERAGE_FILE)\033[0m"
test:	## install and run tests
	make test-install && make test-run
godocs: ## Browse godocs for project using local http server
	@echo "NOTE: Docs will be hosted on http://127.0.0.1:6060"
	cd $(MODULE_PATH);godoc -http=127.0.0.1:6060;cd -