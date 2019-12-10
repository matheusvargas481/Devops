#!/bin/bash

function bake {
    packer build kafka.json
}

function run {
    docker run -d kafka:1
}

function status {
   if docker exec -it "$(docker ps --filter ancestor=kafka:1 -q)" ps aux | grep kafka 
   then
        echo "Kafka Running !"
   else
        echo "Kafka Stop !"
   fi
}

function stop {
    docker stop "$(docker ps --filter ancestor=kafka:1 -q)"
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