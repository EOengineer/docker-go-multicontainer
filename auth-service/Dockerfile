From golang:latest

ADD . /go/src/github.com/eoengineer/auth-service

RUN go install github.com/eoengineer/auth-service

ENTRYPOINT /go/bin/auth-service

EXPOSE 8080
