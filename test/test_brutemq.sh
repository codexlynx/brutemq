#!/usr/bin/env bash

PASSWORD="supertest"

CONTAINER_ID=$(docker run -d --name rabbitmq -p 5672:5672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=${PASSWORD} rabbitmq:alpine)
cd ..

for pass in {1..100}; do
  uuidgen  | md5sum | cut -d'-' -f1 >> pass.txt
done
echo ${PASSWORD} >> pass.txt

./brutemq -u admin -f pass.txt

docker stop ${CONTAINER_ID}
