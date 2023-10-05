# go-playground
A playground for Golang experimentation

## Running in Docker

* `docker build -t go-playground .` - Build the Docker image
* `docker run go-playground` - Run without arguments<br/>
  OR <br/>
  `docker run go-playground /app/go-playground <args>` - Run with the given arguments

## Running Locally

* `go run ./main <args>` - Run the main.go file with the given arguments

## Arguments

| Argument         | Description               |
|------------------|---------------------------|
| `-h` \| `--help` | Print the help message    |
| `-lt`            | Log with timestamps (UTC) |