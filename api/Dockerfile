FROM golang:1.10.3

RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir /go/src
ADD ./api /go/src

WORKDIR /go/src
