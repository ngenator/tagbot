#!/usr/bin/env bash

protoc -I dice/ dice/dice.proto --go_out=plugins=grpc:dice