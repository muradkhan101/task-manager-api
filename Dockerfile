FROM golang:alpine
COPY . /go/src/github.com/task-manager-api
WORKDIR /go/src/github.com/task-manager-api/cmd/server

RUN apk update && apk add git

RUN go get -d -v
RUN go build -o ../../main

EXPOSE 3000
CMD ["/go/src/github.com/task-manager-api/main"]