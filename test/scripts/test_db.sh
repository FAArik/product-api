#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest

echo "Postgresql starting"
sleep 3

docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"
sleep 3 
echo "Database productapp created"

docker exec -it postgres-test psql -u postgres -d productapp -c "
create table if not exist products 
(
    id bigserial not null primary key,  
    name varchar(255) not null,
    price double precision not null,
    discount double precision,
    store varchar(255) not null
);
"
echo "Table products created"
