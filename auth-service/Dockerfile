From golang:latest

ADD . /go/src/github.com/eoengineer/auth-service

COPY ./rerun-scripts/auth /go/bin

RUN go install github.com/eoengineer/auth-service

ENTRYPOINT /go/bin/init

CMD ["bin/auth"]

EXPOSE 8080
