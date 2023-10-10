# go-read-var-log

REST `/var/log` reader

## Contents

* [Assumptions](#assumptions)
* [Endpoints](#endpoints)
* [Enabling Tech Stack](#enabling-tech-stack)
* [Design](#design)
* [Generation of Sample Input Log Files](#generation-of-sample-input-log-files)
* [Arguments](#arguments)
* [Running in Docker](#running-in-docker)
* [Running Locally](#running-locally)
* [Run Unit Tests Locally](#run-unit-tests-locally)
* [Sample Output](#sample-output)
  * [Get list of all log files in /var/log](#get-list-of-all-log-files-in-varlog)
  * [Query the contents of a 1GB file with 9.2 million logs](#query-the-contents-of-a-1gb-file-with-92-million-logs)

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

## Arguments

| Argument         | Description                            |
|------------------|----------------------------------------|
| `-h` \| `--help` | Print the help message                 |
| `-lt`            | Log with timestamps (UTC)              |
| `-n`             | Number of lines to return (default 25) |
| `-p`             | HTTP Port to listen on (default 8080)  |

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

## Sample Output

### Get list of all log files in `/var/log`

```bash
❯ curl "localhost/api/v1/logs"
1GB-9million.log
daily.out
displaypolicyd.stdout.log
fct_uninstall.log
fctinstallpost.log
fsck_apfs.log
fsck_apfs_error.log
fsck_hfs.log
install.log
monthly.out
output.log
shutdown_monitor.log
system.log
weekly.out
wifi.log
```

### Query the contents of a 1GB file with 9.2 million logs

`/var/log/1GB-9million.log` was generated with [GitHub mikebd/go-make-log](https://github.com/mikebd/go-make-log).

```bash
❯ go run ./main -p 80
...
logs.go:30: /api/v1/logs/1GB-9million.log?q=error&r=\sfecig$ returned 1 lines in 2.75945789s
logs.go:30: /api/v1/logs/1GB-9million.log?q=error&r=\sfecig$ returned 1 lines in 2.534780081s
logs.go:30: /api/v1/logs/1GB-9million.log?q=error&r=\sfecig$ returned 1 lines in 2.148826909s
logs.go:30: /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$ returned 1 lines in 1.427393941s
logs.go:30: /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$ returned 1 lines in 1.447307855s
logs.go:30: /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$ returned 1 lines in 1.421464319s
```

☝️ was generated by the following `curl` commands:

```bash
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=error&r=\sfecig$" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=error&r=\sfecig$" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=error&r=\sfecig$" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$" 
```

The `curl` output for all of the above is:

```bash
2023-10-10T09:27:32.029729Z|error|tesak awehit atep itego anetar utiyav oset h om tulad palopi ac arinem eralo fecig
```