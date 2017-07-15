FROM golang:1.9

RUN go get github.com/bwmarrin/discordgo
RUN go get github.com/zalfonse/tagbot

WORKDIR $GOPATH/src/github.com/zalfonse/tagbot
RUN go build

ENTRYPOINT ["./tagbot"]
