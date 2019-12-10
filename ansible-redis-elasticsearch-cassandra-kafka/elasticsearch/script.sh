#!/bin/bash

function bake {
    packer build elasticsearch.json
}

function run {
    docker run -d -p 8080:9200 elasticsearch:1
}

function status {
   curl localhost:8080
}

function stop {
    docker stop "$(docker ps --filter ancestor=elasticsearch:1 -q)"
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