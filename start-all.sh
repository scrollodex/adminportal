#!/bin/bash

DNAME="${1:-myportal}"
NNAME="${1:-net$DNAME}"

set -e

docker container prune -f 

docker build -t scrollodex-adminportal .

docker network create myNetwork || true

docker container stop /myredis || true
docker run --name myredis \
  --network myNetwork \
  -p 6379:6379 \
  -d redis redis-server \
  --save 60 1 \
  --loglevel warning
docker container stop /"$DNAME" || true
docker run --name "$DNAME" \
  --env-file .env \
  --network myNetwork \
  -p 3000:3000 \
  -v ~/gitthings/scrollodex-db-bi:/Users/tlimoncelli/gitthings/scrollodex-db-bi \
  -v ~/gitthings/scrollodex-db-poly:/Users/tlimoncelli/gitthings/scrollodex-db-poly \
  -v /tmp/sessions:/tmp/sessions \
  -it scrollodex-adminportal
