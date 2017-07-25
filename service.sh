#!/bin/bash

command=$1
service=$2

function start {
  pushd services/$service > /dev/null
    docker build -t local/$service .
  popd > /dev/null
  docker run -d --network=tagbot --name=$service local/$service
}

function stop {
  docker rm -f $service
}

function restart {
  stop
  start
}

function usage {
  echo "Usage: ./service {start|stop|restart} {service_name}"
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
