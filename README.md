# go-read-var-log

REST `/var/log` reader

## Assumptions

* The log files are in `/var/log`
* The log files are uncompressed plain text files
  * Compressed or archived files (`.gz`, `bz2`, `.tar`, `.tgz`) are not supported
* The log files are sorted by timestamp (oldest first) as it is captured in each log event
* Log events are separated by a newline character
* REST response Content-Type is text/plain; charset=utf-8, regardless of the Accept request header

## Endpoints

* `GET /api/v1/logs` - Get list of all log files in `/var/log` that this service can read
* `GET /api/v1/logs/{log}` - Get the contents of the log file specified by `{log}` in `/var/log` (e.g. `/api/v1/logs/messages`)
  <br/>Query Parameters:
  * `n` - Number of lines to return (default: 25)
  * `q` - Case sensitive fixed text filter to apply to each line (default: `""`), more performant than `r`
  * `r` - Regex filter to apply to each line (default: `.*`)

## Enabling Tech Stack

* [Go](https://golang.org/) - Programming language
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
* `docker run go-read-var-log` - Run without arguments<br/>
  OR <br/>
  `docker run go-read-var-log /app/go-read-var-log <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Arguments

| Argument         | Description                           |
|------------------|---------------------------------------|
| `-h` \| `--help` | Print the help message                |
| `-lt`            | Log with timestamps (UTC)             |
| `-p`             | HTTP Port to listen on (default 8080) |