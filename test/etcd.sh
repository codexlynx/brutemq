#!/usr/bin/env bash

export PASSWORD="supertest"

CONTAINER_ID=$(docker run -d --rm -p 2379:2379 \
  --name etcd quay.io/coreos/etcd:latest etcd \
  --initial-advertise-peer-urls http://0.0.0.0:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://0.0.0.0:2380 \
  --name node1)

echo -n > passwords.txt
for pass in {1..100}; do
  uuidgen  | md5sum | cut -d'-' -f1 | sed 's| ||g' >> passwords.txt
done
echo "${PASSWORD}" >> passwords.txt

sleep 20
docker run -it --net=host -e ETCDCTL_API=3 quay.io/coreos/etcd etcdctl --endpoints http://0.0.0.0:2379 user add "root:${PASSWORD}"
docker run -it --net=host -e ETCDCTL_API=3 quay.io/coreos/etcd etcdctl --endpoints http://0.0.0.0:2379 auth enable --user "root:${PASSWORD}"
./dist/brutemq_amd64 etcd -d passwords.txt

echo
docker stop "${CONTAINER_ID}"
echo
