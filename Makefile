build:
	go build -o ./bin/server ./src
start:
	./bin/server
dev:
	air
format:
	gofmt -w .
test:
	PARALLEL_CONVEY=false make test-lightspeed
test-lightspeed:
	go test -v --count=1 ./tests/...
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...
install:
	go install github.com/air-verse/air@v1.61.7
	@echo "\033[0;32mAir installed successfully. You can now run 'make dev' to start the development server.\033[0m"
	go install github.com/evilmartians/lefthook@v1.11.12
	lefthook install
	@echo "\033[0;32mLefthook installed and configured successfully.\033[0m"
	@which npm > /dev/null && \
		npm install -g @commitlint/config-conventional@17.6.5 @commitlint/cli@17.6.5 && \
		echo "\033[0;32mCommitlint installed successfully.\033[0m" || \
		echo "\033[0;31mNode is not installed. Please install Node.js to use commitlint.\033[0m"
	go mod tidy
	@echo "\033[0;32mGo modules installed successfully.\033[0m"
sandbox:
	docker compose -f ./infrastructure/docker-compose.yml up
teardown:
	docker compose -f ./infrastructure/docker-compose.yml down
