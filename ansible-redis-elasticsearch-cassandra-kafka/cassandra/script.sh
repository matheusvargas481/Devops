#!/bin/bash

function bake {
    packer build cassandra.json
}

function run {
    docker run -d cassandra:1
}

function status {
    docker exec -it "$(docker ps --filter ancestor=cassandra:1 -q)" /cassandra/apache-cassandra-3.11.5/bin/nodetool status
}

function stop {
    docker stop "$(docker ps --filter ancestor=cassandra:1 -q)"
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