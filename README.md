# go-read-var-log

REST `/var/log` reader

## Assumptions

Some assumptions could be considered TODO items for future enhancement

* The log files are in `/var/log`
* The log files are uncompressed plain text files
  * Compressed or archived files (e.g. `.gz`, `bz2`, `.tar`, `.tgz`) are not supported
* The log files are sorted by timestamp (oldest first) as it is captured in each log event
* Log events are separated by a newline character
* REST response Content-Type is text/plain; charset=utf-8, regardless of the Accept request header
* All errors are returned to REST callers as HTTP status 500, even if these might be correctable by the caller
* LRU result caching of compiled regexes and search results are not implemented
* Testing coverage should be added for the `controller` package

## Endpoints

* `GET /api/v1/logs` - Get list of all log files in `/var/log` that this service can read
* `GET /api/v1/logs/{log}` - Get the contents of the log file specified by `{log}` in `/var/log` (e.g. `/api/v1/logs/system.log`)
  <br/>Query Parameters:
  * `n` - Number of lines to return (default: 25, all: 0)
  * `q` - Case sensitive fixed text filter to apply to each line (default: `""`), more performant than `r`
  * `r` - Regex filter to apply to each line (default: `.*`)
  * Note: `q` and `r` may both be specified, but at most once each.  Multiple `q` or `r` parameters are ignored.

## Enabling Tech Stack

* [Go](https://golang.org/) - Programming language
  * If not installed, consider: `brew install golang` or similar as appropriate for 
    how you manage packages for development on your OS
  * Installation is not required if running in Docker
* [HttpRouter](https://github.com/julienschmidt/httprouter) - HTTP request router

## Design

* The `service` package contains the logic of the service, agnostic of presentation and transport
* The `controller` package contains the HTTP request handlers

## Generation of Sample Input Log Files

Optionally, arbitrarily large fake log files can be generated in json or delimited format to augment any that
are naturally in `/var/log`.

* See: [GitHub mikebd/go-make-log](https://github.com/mikebd/go-make-log) - Generate fake application log files (e.g. as input to log processors)

## Running in Docker

* `docker build -t go-read-var-log .` - Build the Docker image
* `docker run -v /var/log:/var/log:ro -p 80:8080 go-read-var-log <args>` - Run with the given arguments
  <br/>Prevents writing to `/var/log`.  When run without `<args>`, this enables the following from the docker host:
  * `curl "localhost/api/v1/logs"`
  * `curl "localhost/api/v1/logs/system.log?n=10&q=syslog"`

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Run Unit Tests Locally

Tests are implemented in `*_test.go` files in the same package as the code under test.
Test data files are stored under `testdata` directories, which are ignored by `go build` and `go run`.

* `go test ./... -count=-1` - Run all unit tests
* `go test ./... -count=-1 -cover` - Run all unit tests with coverage
  * Shows 83% coverage of the `service` package 

## Arguments

| Argument         | Description                            |
|------------------|----------------------------------------|
| `-h` \| `--help` | Print the help message                 |
| `-lt`            | Log with timestamps (UTC)              |
| `-n`             | Number of lines to return (default 25) |
| `-p`             | HTTP Port to listen on (default 8080)  |