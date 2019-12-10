#!/bin/bash

function bake {
    packer build redis.json
}

function run {
    docker run -d redis:1
}

function status {
   if docker exec -it "$(docker ps --filter ancestor=redis:1 -q)" ps aux | grep redis 
   then
        echo "Redis Running !"
   else
        echo "Redis Stop !"
   fi
}

function stop {
    docker stop "$(docker ps --filter ancestor=redis:1 -q)"
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