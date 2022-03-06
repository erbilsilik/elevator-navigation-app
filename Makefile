APP=elevator-navigation-app

DIRECTORIES = $(shell go list ./... | grep pkg)

run: ## Runs the applications
	go mod tidy
	go build
	chmod +x ./elevator-navigation-app
	./elevator-navigation-app

test-unit: ## Runs unit tests
		go test $(DIRECTORIES) -v -count=1  | tee report/unit_test.out

create-report: ## Creates reports directory for coverage outputs
	mkdir -p report
coverage: ## Creates coverage report
	go test $(DIRECTORIES) -json -cover -coverprofile=report/coverage.out
coverage-html: ## Creates coverage html from report
	go tool cover -html=report/coverage.out

test: create-report test-unit coverage coverage-html

help: ## Displays this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'