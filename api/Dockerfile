FROM golang:1.10.3

RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir /go/src/app
ADD ./app /go/src/app
ADD ./startserver.sh /go/src/app/

WORKDIR /go/src/app
