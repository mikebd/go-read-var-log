FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD config ./config
ADD controller ./controller
ADD main ./main
ADD service ./service

RUN go build -o go-read-var-log ./main

ENTRYPOINT ["./go-read-var-log"]
