FROM golang:1.23-alpine

RUN go install github.com/githubnemo/CompileDaemon@1.3.0

WORKDIR /service

COPY ./ /service

ENTRYPOINT CompileDaemon \
    -build="go build -o /service/main /service/cmd/main.go" \
    -command="/service/main" \
    -directory="cmd" \
    -directory="service" \
    -directory="explore" \
    -exclude-dir=".git"
