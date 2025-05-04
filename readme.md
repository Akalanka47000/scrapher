# Scrapher Service

Go service for the Scrapher web application. 

## [API Documentation](https://documenter.getpostman.com/view/32343835/2sB2j4grCE)

## Prerequisites
 - [Node](https://nodejs.org/en) - If you want to make use of [commitlint](https://commitlint.js.org) (optional)

## Getting started

- Run `make install` to download all dependencies and install the required tools. This is required only once. Afterwards you could use the traditional `go mod tidy` for dependency management.
- Run `make dev` to start a development server with hot reloading (using Air).
- Run `make test` to run the tests.
- Run `make test-coverage` to run the tests and generate a coverage report.
- Run `make lint` to run the linter.
- Run `make build` to build the application.