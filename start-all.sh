#!/bin/bash

DNAME="${1:-myportal}"
NNAME="${1:-net$DNAME}"

docker container prune -f 

docker build -t scrollodex-adminportal .

docker network create myNetwork

docker container stop /myredis
docker run --name myredis \
  --network myNetwork \
  -p 6379:6379 \
  -d redis redis-server \
  --save 60 1 \
  --loglevel warning
docker container stop /"$DNAME"
docker run --name "$DNAME" \
  --env-file .env \
  --network myNetwork \
  -p 3000:3000 \
  -v ~/gitthings/scrollodex-db-bi:/Users/tlimoncelli/gitthings/scrollodex-db-bi \
  -v ~/gitthings/scrollodex-db-poly:/Users/tlimoncelli/gitthings/scrollodex-db-poly \
  -it scrollodex-adminportal
