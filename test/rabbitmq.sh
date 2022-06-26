#!/usr/bin/env bash

export PASSWORD="supertest"

CONTAINER_ID=$(docker run -d --rm -p 5672:5672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=${PASSWORD} rabbitmq:alpine)

echo > passwords.txt
for pass in {1..100}; do
  uuidgen  | md5sum | cut -d'-' -f1 >> passwords.txt
done
echo ${PASSWORD} >> passwords.txt

sleep 20
./dist/brutemq_amd64 amqp -u admin -d passwords.txt -e localhost:5672/

echo
docker stop "${CONTAINER_ID}"
