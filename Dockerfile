FROM golang

ADD . /go/src/github.com/cormierm/TvtvToXmlTV

RUN go install github.com/cormierm/TvtvToXmlTV

ENTRYPOINT /go/bin/TvtvToXmlTV

EXPOSE 8080