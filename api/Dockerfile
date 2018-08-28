FROM golang:1.10.3

RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/api
ADD . /go/src/api

WORKDIR /go/src/api

RUN dep ensure
