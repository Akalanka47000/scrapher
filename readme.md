# Scrapher Service

Go service for the Scrapher web application. 

## [API Documentation](https://documenter.getpostman.com/view/32343835/2sB2j4grCE)

## Prerequisites
 - [Go 1.24](https://golang.org/dl) - The Go programming language
 - [Chrome](https://www.google.com/chrome) - Used under the hood for headless browser automation by [go-rod](https://go-rod.github.io)
 - [Node (optional)](https://nodejs.org/en) - If you want to make use of [commitlint](https://commitlint.js.org)
 - [Docker (optional)](https://www.docker.com) - If you want to check out the complete setup with monitoring and logging 

## Getting started

- Run `make install` to download all dependencies and install the required tools. This is required only once. Afterwards you could use the traditional `go mod tidy` for dependency management.
- Run `make dev` to start a development server with hot reloading (using Air).
- Run `make test` to run all tests suites.
- Run `make test-lightspeed` to run the same above tests cost faster at the cost of readability.
- Run `make lint` to run the linter.
- Run `make build` to build the application.
- Run `make start` to start the built application (You need to run `make build` first).

#### Environment variables are loaded from the `.env` file. You can create one by copying the `.env.example` file. The server won't start if any of the required environment variables are not set or invalid, an error message will tell you more about it.

# Full setup with client, monitoring and logging

- Run `make sandbox` to spin up a complete setup with the client, monitoring and logging. 

    #### Client application will be available at [http://localhost:5173](http://localhost:5173)
    #### Scrapher service will be available at [http://localhost:8080](http://localhost:8080)
    #### Grafana will be available at [http://localhost:3000](http://localhost:3000)

- Under the hood this will start the following services:
  - [Grafana](https://grafana.com) - Explore logs and metrics. There already is a dashboard preconfigured for metrics.
  - [Prometheus](https://prometheus.io) - Collect metrics from the application and allow Grafana to query them
  - [Loki](https://grafana.com/oss/loki) - Log aggregator
  - [Promtail](https://grafana.com/docs/loki/latest/clients/promtail/) - Collect logs from the application and send them to Loki

- Run `make teardown` to stop all services started by `make sandbox` when you are done testing.

## Directory structure

- .github - GitHub workflows
-  docs - Contains the Postman API Specification, automatically generated and commited by the [Postman Github Integration](https://learning.postman.com/docs/integrations/available-integrations/github/)
- infrastructure - Contains the Docker Compose file and other infrastructure related files
   - grafana - Contains the Grafana configuration
      - dashboards - Contains preconfigured dashboards
      - datasources - Contains the data source configuration

- src
   - config - Environmental configuration
   - global - Global packages
   - middleware - Custom request/response interceptors
   - modules - Contains the application modules
        - analysis - Contains the analysis module
           - api - Contains the API handlers
            - v1 - Version 1 of the API
               - dto - Data Transfer Objects
  - pkg - Custom extensions of external librarie which can be shared between modules
  - utils - Utility functions
  - app.go - The fiber application which sets up almost everything
  - server.go - The main entry point of the application

- test - Contains all test suites
  - __mocks__ - Contains mock files
  - integration - Contains integration tests
  - performace - Contains Jmeter scripts used for stress and load testing
  - unit - Contains unit tests

## What's special about this project?

This project is designed for scalability and performance. It incorporates a lot of best practices and patterns to ensure robustness and maintainability. Some of the key features include:
- **Modular architecture which is extremely useful when the API is growing and you need to add new features. Each module has its own API handlers, DTOs, and services**
- **Structured logging with everything you need for traceability**
- **An opinionated way of handling errors through the use of panics and a global error handler. Keeps the code extremely clean and readable**
- **Proper health checks and graceful shutdown**
- **DX focused enhanced request validation built on top of [go-playground/validator](https://github.com/go-playground/validator)**
- **Custom middlware such as rate limiting, request logging, and response caching**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.