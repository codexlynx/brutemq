#!/usr/bin/env bash

export PASSWORD="supertest"

CONTAINER_ID=$(docker run -d --rm --name rabbitmq -p 5672:5672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=${PASSWORD} rabbitmq:alpine)

echo > pass.txt
for pass in {1..100}; do
  uuidgen  | md5sum | cut -d'-' -f1 >> pass.txt
done
echo ${PASSWORD} >> pass.txt

sleep 20
./brutemq -u admin -f pass.txt --url localhost:5672/

echo
docker stop ${CONTAINER_ID}
