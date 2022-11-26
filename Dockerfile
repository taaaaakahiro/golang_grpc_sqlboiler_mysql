FROM golang:1.19.1-alpine

RUN apk update && apk add git

WORKDIR /go/src

ADD . /go/src

CMD [ "go", "run", "./pkg/server/main.go" ]

