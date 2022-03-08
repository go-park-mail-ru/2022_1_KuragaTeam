#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    create table if not exists users
    (
        id serial constraint users_pk primary key,
        username varchar(50)  not null,
        email    varchar(50)  not null,
        password varchar(250) not null,
        salt     varchar(50)  not null
    );

    create unique index users_email_uindex
        on users (email);
  COMMIT;
EOSQL