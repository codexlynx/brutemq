#!/usr/bin/env bash

export PASSWORD="supertest"

CONTAINER_ID=$(docker run -d --rm -p 8200:8200 --name vault -e VAULT_DEV_ROOT_TOKEN_ID=root vault:latest)

echo -n > passwords.txt
for pass in {1..100}; do
  uuidgen  | md5sum | cut -d'-' -f1 | sed 's| ||g' >> passwords.txt
done
echo "${PASSWORD}" >> passwords.txt

sleep 20

docker run -it --net=host -e VAULT_ADDR=http://0.0.0.0:8200 -e VAULT_TOKEN=root vault auth enable userpass
docker run -it --net=host -e VAULT_ADDR=http://0.0.0.0:8200 -e VAULT_TOKEN=root vault write auth/userpass/users/root password=${PASSWORD} policies=root
./dist/brutemq_amd64 vault -d passwords.txt

echo
docker stop "${CONTAINER_ID}"
echo
