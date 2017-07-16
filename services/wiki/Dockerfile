FROM golang:1.9

ADD . $GOPATH/src/github.com/zalfonse/wiki

WORKDIR $GOPATH/src/github.com/zalfonse/wiki
RUN go build

ENTRYPOINT ["./wiki"]
