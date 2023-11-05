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
* [Benchmark with Apache Bench](#benchmark-with-apache-bench)

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
  * Search result caching would require invalidation when the file changes
  * Hot searches could have dedicated caches that are eagerly refreshed when the file changes
* Testing coverage should be added for the `controller` package
* Add `GOMAXPROCS` to Dockerfile or set that in the code
* There is a [get-large-log](https://github.com/mikebd/go-read-var-log/tree/get-large-log) branch where I hope to
  experiment with a more efficient way to read very large files, but it is not yet complete and is not formally
  part of this submission

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
| `-l`             | Large File Bytes (default 10,000,000)  |
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
* <b>IMPORTANT</b>: Improve scalability with GOMAXPROCS: `GOMAXPROCS=8 go run ./main <args>`

## Run Unit Tests Locally

Tests are implemented in `*_test.go` files in the same package as the code under test.
Test data files are stored under `testdata` directories, which are ignored by `go build` and `go run`.

* `go test ./... -count=-1` - Run all unit tests
* `go test ./... -count=-1 -cover` - Run all unit tests with coverage

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
logs.go:31: /api/v1/logs/1GB-9million.log?q=error&r=\sfecig$&n=1 returned 1 lines using large strategy in 1.23877778s
logs.go:31: /api/v1/logs/1GB-9million.log?q=error&r=\sfecig$&n=1 returned 1 lines using large strategy in 1.202586854s
logs.go:31: /api/v1/logs/1GB-9million.log?q=error&r=\sfecig$&n=1 returned 1 lines using large strategy in 1.198800304s
logs.go:31: /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1 returned 1 lines using large strategy in 856.594448ms
logs.go:31: /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1 returned 1 lines using large strategy in 866.171423ms
logs.go:31: /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1 returned 1 lines using large strategy in 863.319268ms
```

☝️ was generated by the following `curl` commands:

```bash
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=error&r=\sfecig$&n=1" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=error&r=\sfecig$&n=1" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=error&r=\sfecig$&n=1" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1" 
❯ curl "localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1" 
```

The `curl` output for all of the above is:

```bash
2023-10-10T09:27:32.029729Z|error|tesak awehit atep itego anetar utiyav oset h om tulad palopi ac arinem eralo fecig
```

## Benchmark with Apache Bench

[Apache Bench](https://httpd.apache.org/docs/2.4/programs/ab.html) is a simple HTTP load testing tool.

### Small file

`/var/log/daily.out` is a 498KB file with 7,992 lines.

```bash
❯ ab -c 6 -k -n 10000 'localhost/api/v1/logs/daily.out'

Document Path:          /api/v1/logs/daily.out
Document Length:        1930 bytes

Concurrency Level:      6
Time taken for tests:   2.062 seconds
Complete requests:      10000
Failed requests:        0
Keep-Alive requests:    10000
Total transferred:      20730000 bytes
HTML transferred:       19300000 bytes
Requests per second:    4849.47 [#/sec] (mean)
Time per request:       1.237 [ms] (mean)
Time per request:       0.206 [ms] (mean, across all concurrent requests)
Transfer rate:          9817.34 [Kbytes/sec] received

Connection Times (ms)
min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.3      1       3
Waiting:        0    1   0.3      1       3
Total:          0    1   0.3      1       3

Percentage of the requests served within a certain time (ms)
50%      1
66%      1
75%      1
80%      1
90%      2
95%      2
98%      2
99%      2
100%     3 (longest request)
```

### Large file with fixed text and regex matching

#### Non-default GOMAXPROCS=8

```bash
❯ ab -c 6 -k -n 1000 'localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1'

Document Path:          /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1
Document Length:        117 bytes

Concurrency Level:      6
Time taken for tests:   259.356 seconds
Complete requests:      1000
Failed requests:        0
Keep-Alive requests:    1000
Total transferred:      259000 bytes
HTML transferred:       117000 bytes
Requests per second:    3.86 [#/sec] (mean)
Time per request:       1556.135 [ms] (mean)
Time per request:       259.356 [ms] (mean, across all concurrent requests)
Transfer rate:          0.98 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   909 1550  51.8   1541    1788
Waiting:      908 1550  51.8   1540    1788
Total:        909 1550  51.8   1541    1788

Percentage of the requests served within a certain time (ms)
  50%   1541
  66%   1559
  75%   1568
  80%   1578
  90%   1605
  95%   1632
  98%   1692
  99%   1757
 100%   1788 (longest request)
```

#### Default GOMAXPROCS=1

```bash
❯ ab -c 6 -k -n 1000 'localhost/api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1'

Document Path:          /api/v1/logs/1GB-9million.log?q=|error|&r=\sfecig$&n=1
Document Length:        117 bytes

Concurrency Level:      6
Time taken for tests:   279.119 seconds
Complete requests:      1000
Failed requests:        0
Keep-Alive requests:    1000
Total transferred:      259000 bytes
HTML transferred:       117000 bytes
Requests per second:    3.58 [#/sec] (mean)
Time per request:       1674.712 [ms] (mean)
Time per request:       279.119 [ms] (mean, across all concurrent requests)
Transfer rate:          0.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:  1276 1663  53.5   1654    2018
Waiting:     1276 1663  53.5   1653    2018
Total:       1276 1663  53.5   1654    2018

Percentage of the requests served within a certain time (ms)
  50%   1654
  66%   1667
  75%   1679
  80%   1687
  90%   1714
  95%   1745
  98%   1803
  99%   1932
 100%   2018 (longest request)
```
