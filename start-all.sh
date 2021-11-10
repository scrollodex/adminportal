#!/bin/bash

DNAME="${1:-myportal}"
NNAME="${1:-net$DNAME}"

docker container prune -f 

docker build -t scrollodex-adminportal .

docker network create myNetwork

docker container stop /myredis
docker run --network myNetwork -p 6379:6379 --name myredis -d redis redis-server --save 60 1 --loglevel warning
docker container stop /"$DNAME"
docker run --network myNetwork --env-file .env -p 3000:3000 --name "$DNAME" -it scrollodex-adminportal
