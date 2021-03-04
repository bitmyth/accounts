FROM golang:1.15-alpine

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git gcc musl-dev\
    && rm -rf /var/cache/apk/*

RUN go get -u github.com/pressly/goose/cmd/goose
