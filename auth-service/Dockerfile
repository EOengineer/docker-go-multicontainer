FROM gliderlabs/alpine:3.3

RUN apk add git --update
RUN apk add go --update

ENV GOPATH /go
ENV PATH $GOPATH/BIN:$PATH

# ADD . /go/src/github.com/eoengineer/auth-service

RUN go get github.com/ddollar/init
RUN go get github.com/cleblanc87/rerun

WORKDIR /go/src/github.com/eoengineer/auth-service/
COPY . /go/src/github.com/eoengineer/auth-service/
RUN go install ./...

ENTRYPOINT ["/go/bin/init"]

CMD ["./bin/auth-service"]

EXPOSE 8080
