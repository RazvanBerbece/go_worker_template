# Project Title

A repository template for containerised backend workers written in the Go language. 

Comes with:
- containerised MySQL image
- 1-command service composition for all services, external and Go alike: [docker-compose.yml](./docker-compose.yml)
- CI pipeline to run in GitHub: [GHA CI Workflow](./.github/workflows/ci.yml)
- helper Go functions and structs for DB connections and queries

## Getting Started

### Prerequisites

* Describe any prerequisites, libraries, OS version, services, etc., needed before installing and running program.
* ex. Windows 10

### How to Run

1. Run `docker compose up -d --remove-orphans --build` in the root of the repo to start the
  * MySQL instance
  * Go worker service

2. Run `docker compose down -v` to bring down all infrastructure and stop all the services

### Run tests
The automated test suite can be run through the commands below, executed in the root of this repository.

```
docker compose down -v && docker-compose -f docker-compose.ci.yml up --remove-orphans --force-recreate --build --exit-code-from sut
```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```