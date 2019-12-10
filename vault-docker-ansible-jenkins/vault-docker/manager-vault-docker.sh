#!/bin/bash

function bake {
   docker build -t vault-docker:1 .
}

function run {
    docker run -d -p 8200:8200 vault-docker:1 vault-docker:1
}

function status {
   docker exec -it "$(docker ps --filter ancestor=vault-docker:1 -q)" vault status
}

function stop {
    docker stop "$(docker ps --filter ancestor=vault-docker:1 -q)"
}

case $1 in
    "bake")
        bake
        ;;
    "run")
        run
        ;;
    "status")
        status
        ;;
    "stop")
        stop
        ;;
    *)
esac