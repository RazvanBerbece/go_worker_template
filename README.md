# Project Title

A repository template for containerised backend workers written in the Go language.

Comes with:
- containerised MySQL image
- 1-command service composition for all services, external and Go alike: [docker-compose.yml](./docker-compose.yml)
- CI pipeline to run in GitHub: [GHA CI Workflow](./.github/workflows/ci.yml)
- helper Go functions and structs for the DB layer using [Gorm](https://github.com/go-gorm/gorm)

## Getting Started

### Prerequisites

* Go
* Docker
* docker compose

### How to Run

1. Run `docker compose up -d --remove-orphans --build` in the root of the repo to start the
  * MySQL instance
  * Go worker service

2. Run `docker compose down -v` to bring down all infrastructure and stop all the services

### Run the tests
The automated test suite can be run through the command below, executed in the root of this repository.

```
docker-compose -f docker-compose.ci.yml up --remove-orphans --force-recreate --build --exit-code-from sut
```