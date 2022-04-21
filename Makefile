AWS_PROFILE=default
CSV_FILE=${CSV_FILE_50_ITEM}
RESOURCE_URL=${RESOURCE_URL_LOCAL}

CSV_FILE_1_ITEM=extra/data/the_movie_database/movies_metadata_oneitem.csv
CSV_FILE_50_ITEM=extra/data/the_movie_database/movies_metadata_small.csv
CSV_FILE_ALL_ITEM=extra/data/the_movie_database/movies_metadata.csv
RESOURCE_URL_LOCAL=https://aezzlrhy4h.execute-api.us-east-1.amazonaws.com/movies
RESOURCE_URL_DEV=https://j7grpr203l.execute-api.us-east-1.amazonaws.com/movies

LINT_PATH=reports/lint
LINT_FILE=lint_report_$(shell date '+%Y-%m-%d-%H%M%S').html




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
build:	## Start "local" environment with supporting cloud resources
	AWS_PROFILE=${AWS_PROFILE} npx sst build
start:	## Alias for "deploy-local"
	make deploy-local
deploy-local:	## Start "local" environment with supporting cloud resources, NOTE: use "local" stage prefix
	AWS_PROFILE=${AWS_PROFILE} npx sst start --stage local
deploy:	## Deploy stack as complete build to AWS environment, NOTE: use "dev" stage prefix
	AWS_PROFILE=${AWS_PROFILE} npx sst deploy --stage dev
remove-env-local:	## Remove "local" environment and supporing cloud resources
	AWS_PROFILE=${AWS_PROFILE} npx sst remove --stage local
remove-env:	## Remove "local" environment
	AWS_PROFILE=${AWS_PROFILE} npx sst remove --stage dev

remove-all:	## Remove ALL environments including local
	make remove-env-local
	make remove
	rm -Rf .build
data-import-local:	## Import test data into local environment
	@echo "WARNING: Update ${RESOURCE_URL} if needed."
	go run scripts/cmd/import_data/import_data.go -u=${RESOURCE_URL} -p=${CSV_FILE}
data-import:	## Import test data into production environment
	@echo "WARNING: Update ${RESOURCE_URL} if needed."
	go run scripts/cmd/import_data/import_data.go -u=${RESOURCE_URL} -p=${CSV_FILE}
data-import-fast:	## Fast Import test data into production environment
	@echo "WARNING: Update ${RESOURCE_URL} if needed."
	go run scripts/cmd/import_data/import_data.go -f -u=${RESOURCE_URL} -p=${CSV_FILE}
mod-tidy:	## Tidy up the go module dependencies
	# Clean up any uncessessary go modules in go.mod
	go mod tidy
sst-console:	## View the SST Console
	AWS_PROFILE=${AWS_PROFILE} npx sst console
lint-install: ## install lining executable tool, golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint-run: ## Run lint without the install
	$(eval GP=$(shell go env GOPATH))
	$(eval XP=$(shell echo './node_modules/xunit-viewer/bin'))
	$(GP)/bin/golangci-lint --out-format junit-xml run > tmp.xml; $(XP)/xunit-viewer -r tmp.xml -o $(LINT_PATH)/$(LINT_FILE); rm -rf tmp.xml
	cp -f $(LINT_PATH)/$(LINT_FILE) $(LINT_PATH)/lint_report_latest.html
lint: ## run both lint install and run
	make lint-install && make lint-run
test: ## run tests
	@echo "NOT IMPLEMENTED. Add test command here"
godocs: ## Browse godocs for project using local http server
	@echo "NOTE: Docs will be hosted on http://127.0.0.1:6060"
	godoc -http=127.0.0.1:6060
