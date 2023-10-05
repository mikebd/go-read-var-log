# go-read-var-log

REST `/var/log` reader

## Running in Docker

* `docker build -t go-read-var-log .` - Build the Docker image
* `docker run go-read-var-log` - Run without arguments<br/>
  OR <br/>
  `docker run go-read-var-log /app/go-read-var-log <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Arguments

| Argument         | Description               |
|------------------|---------------------------|
| `-h` \| `--help` | Print the help message    |
| `-lt`            | Log with timestamps (UTC) |