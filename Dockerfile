FROM golang:1.15-alpine
MAINTAINER Conjur Inc.

CMD /bin/bash
EXPOSE 8080

RUN apk add --no-cache bash \
                       gcc \
                       git \
                       jq \
                       less \
                       libc-dev

RUN go get -u github.com/jstemmer/go-junit-report && \
    go get -u github.com/smartystreets/goconvey && \
    go get -u github.com/axw/gocov/gocov && \
    go get -u github.com/AlekSi/gocov-xml

WORKDIR /conjur-api-go

COPY go.mod go.sum ./
RUN go mod download

COPY . .
