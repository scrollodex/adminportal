#!/usr/bin/env bash
docker build -t scrollodex-adminportal .
docker run --env-file .env -p 3000:3000 -it scrollodex-adminportal
