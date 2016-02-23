From golang:latest

ADD . /go/src/github.com/eoengineer/translation-service

COPY ./rerun-scripts/translation /go/bin

RUN go install github.com/eoengineer/translation-service

ENTRYPOINT /go/bin/init

CMD ["bin/translation"]

EXPOSE 8080
