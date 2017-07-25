#!/bin/bash

command=$1

function start {
  docker build -t local/tagbot .
  docker run -d -e BOT_TOKEN=`cat .bot_token` --network=tagbot --name=tagbot local/tagbot
}

function stop {
  docker rm -f tagbot
}

function restart {
  stop
  start
}

function usage {
  echo "Usage: ./bot {start|stop|restart}"
}

case "$command" in
  start)
    start;;
  stop)
    stop;;
  restart)
    restart;;
  help)
    usage;;
  *)
    usage
    exit 1
esac
