FROM golang:1.9

ADD . $GOPATH/src/github.com/zalfonse/dice

WORKDIR $GOPATH/src/github.com/zalfonse/dice
RUN go build

ENTRYPOINT ["./dice"]
