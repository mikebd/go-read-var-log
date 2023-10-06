FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD config ./config
ADD main ./main

RUN go build -o go-read-var-log ./main

CMD ["./go-read-var-log"]
