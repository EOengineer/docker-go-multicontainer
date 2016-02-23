From golang:latest

ENV GOPATH /go
ENV PATH $GOPATH/BIN:$PATH

# ADD . /go/src/github.com/eoengineer/translation-service




RUN go get github.com/ddollar/init
RUN go get github.com/cleblanc87/rerun

WORKDIR /go/src/git.arthrex.io/marketing/translation-service/
COPY . /go/src/git.arthrex.io/marketing/translation-service/
RUN go install ./...

ENTRYPOINT ["/go/bin/init"]

CMD ["bin/api"]

EXPOSE 8080
