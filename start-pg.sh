#!/bin/sh

docker volume create snaily-db-data

docker service create \
 --name spg \
 --publish 5000:5432 \
 --limit-memory 32M \
 --mount type=volume,source=snaily-db-data,destination=/var/lib/postgresql/data \
 -e POSTGRES_PASSWORD=snaily \
 -e POSTGRES_DB=madast \
 -e POSTGRES_USER=snaily postgres
 
docker service update spg --network-add my-services
