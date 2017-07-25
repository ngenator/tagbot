FROM golang:1.9

RUN go get github.com/bwmarrin/discordgo
ADD . $GOPATH/src/github.com/zalfonse/tagbot

WORKDIR $GOPATH/src/github.com/zalfonse/tagbot
RUN go get
RUN go build

ENTRYPOINT ["./tagbot"]
