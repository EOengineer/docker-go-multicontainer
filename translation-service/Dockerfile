From golang:latest

ADD . /go/src/github.com/eoengineer/translation-service

RUN go install github.com/eoengineer/translation-service

ENTRYPOINT /go/bin/translation-service

EXPOSE 8080
