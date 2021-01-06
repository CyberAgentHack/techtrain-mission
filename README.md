# gameapi - TechTrainMISSION
[![Go Report Card](https://goreportcard.com/badge/github.com/task4233/techtrain-mission)](https://goreportcard.com/report/github.com/task4233/techtrain-mission)

## Description
 - [Given Specification](./docs/README.md)
 - [API specification](https://github.com/CyberAgentHack/techtrain-mission/blob/master/api-document.yaml)
 - [SwaggerEdittor](https://editor.swagger.io/)

## Requirements
### Using Docker(Recommend)
 - [Docker](https://docs.docker.com/engine/install/)
 - [Docker Compose](https://docs.docker.com/compose/install/#install-compose)

### Self Hosting
 - [golang](https://golang.org/doc/install)
 - MySQL

## Quick Start
```
git clone https://github.com/task4233/techtrain-mission.git
cd techtrain-mission
cp .env.example .env
make docker-up
```

## Environment Variables
 - **Install `.env` on project root**

```
./
├── .env
├── .env.example
├── docker/
├── docs/
├── gameapi/
├── mysql
├── README.md
├── server
├── tools/
├── ...
(snip)
```

 - see [`.env.example`](./.env.example)

### Required
 - `PORT`: `8080` is default setting.
 - `DB_DATABASE`: database name.
 - `DB_USER`: database user name.
 - `DB_PASSWORD`: database user password.
 - `DB_HOST`: `mysql-server` is default setting on using Docker. In case of self hosting, sets IP address for database.
 - `DB_PORT`: `3306` is default setting. Set port for database.
 - `DB_ROOT_PASSWORD`: set your secret for database.
 - `TZ`: `Asia/Tokyo` is default setting.

## Makefile Commands
 - `make docker-up`: Serve using Docker
 - `make down`: Shutdown Docker Containers
 - `make test`: Test codes
 - `make test-with-coverage`: Test codes and export coverage
 - `make build`: Build codes
 - `make run`: Build codes and run
 - `make lint`: Run linters(goimports, gofmt, golint, gsc, gosec, staticcheck, errcheck, misspell)
