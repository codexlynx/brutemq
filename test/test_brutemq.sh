#!/usr/bin/env bash

CONTAINER_ID=$(docker run -d --name rabbitmq -p 5672:5672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=1234 rabbitmq:alpine)

docker stop ${CONTAINER_ID}
