# dashapi

[![Build Status](https://gitlab.com/nmarley/dashapi/badges/master/build.svg)](https://gitlab.com/nmarley/dashapi/pipelines)
[![Go Report](https://goreportcard.com/badge/gitlab.com/nmarley/dashapi)](https://goreportcard.com/badge/gitlab.com/nmarley/dashapi)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/gitlab.com/nmarley/dashapi)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](LICENSE)

> Database and related API which manages Dash Governance info

## Table of Contents
- [Install](#install)
  - [Dependencies](#dependencies)
- [Usage](#usage)
- [Configuration](#configuration)
  - [Quick start](#quick-start)
- [Maintainer](#maintainer)
- [Contributing](#contributing)
- [License](#license)

## Install

Clone the repo (or install via `go get`) and build the project. A makefile has been included for convenience.

```sh
git clone https://gitlab.com/nmarley/dashapi.git
cd dashapi
make
```

### Dependencies

The dashapi reads and writes governance info to a Postgres database, therefore a running instance of the Postgres database is required as a dependency. The connection is configured via environment variables. See [Configuration](#configuration) for more info.

## Usage

First, copy `.env.example` to `.env` and modify accordingly. Postgres variables need to be configured to point to an accessible, running Postgres instance.

```sh
# config
cp .env.example .env
vi .env #  (edit accordingly)

# run
go run dashapi

# -or-
go build
./dashapi
```

## Configuration

This project uses environment variables for configuration. Variables are read from a `.env` file and can be overwritten by variables defined in the environment or directly passed to the process. See all available settings in [.env.example](.env.example).

### Quick start

A `docker-compose` file is included for testing purposes, which also sets up a Postgres database.

```
cp .env.example .env
vi .env #  (edit accordingly)

docker-compose up
```

To verify:

```
curl -i http://127.0.0.1:7001/health
```

### Generating a JWT

Some routes in the API are only available with authentication.

For these, a JWT token must be sent in the header (see `curl_examples.sh` in this repo). There is currently no authentication table or route, so this must be manually generated.

To generate the JWT token, you can use the [JWT Debugger](https://jwt.io/#debugger-io). Simply visit the site, adjust the payload data accordingly, and in place of `your-256-bit-secret`, use the value of `$JWT_SECRET_KEY` that you set in the .env file (can be any secret string). Then click the "Share JWT" button to retrieve the JWT. This is the value you should send in the header after "Authorization: Bearer ".

An example JWT token looks like:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJUZXN0IFRlc3RlcnNvbiIsInN1YiI6IkpvaG4gRG9udXQiLCJpYXQiOjE1NTE0NjYyMjN9.Z03u0ZogZZ4W2C9E7FgisQxWqp-XsnuS48JAxzRxQ1I
```

*Note that this is just an example and will not work with any production deployment.*

## Maintainer

[@nmarley](https://gitlab.com/nmarley)

## Contributing

Feel free to dive in! [Open an issue](https://gitlab.com/nmarley/dashapi/issues/new) or submit PRs.

## License

[ISC](LICENSE)
