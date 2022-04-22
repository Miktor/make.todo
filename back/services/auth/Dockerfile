# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o build/main main.go" --command=./build/main
EXPOSE 8000