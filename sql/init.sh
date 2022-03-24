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
        salt     varchar(50)  not null,
        avatar varchar(100),
        subscription_expires timestamp
    );

    create unique index users_email_uindex
        on users (email);
  COMMIT;

  CREATE TABLE IF NOT EXISTS movies(
       id serial constraint movies_pk primary key,
       name varchar(60) not null unique,
       name_picture varchar(255) not null,
       year smallint not null,
       duration varchar(25) not null,
       age_limit smallint not null,
       description varchar(1024) not null,
       kinopoisk_rating numeric(2,1) not null,
       tagline varchar(255) not null,
       picture varchar(255) not null,
       video varchar(255) not null,
       trailer varchar(255) not null
  );

  CREATE TABLE genre(
     id serial constraint genre_pk primary key,
     name varchar(255) unique
  );

  BEGIN;
  CREATE TABLE IF NOT EXISTS movies_genre(
      id serial constraint movies_genre_pk primary key,
      movie_id int,
      genre_id int
  );

  create unique index movies_genre_uindex
      on movies_genre (movie_id, genre_id);
  COMMIT;

  CREATE TABLE country(
      id smallserial constraint country_pk primary key,
      name varchar(255) unique
  );

  BEGIN;
  CREATE TABLE movies_countries(
      id serial constraint movies_countries_pk primary key,
      movie_id int,
      country_id int
  );

  create unique index movies_countries_uindex
      on movies_countries (movie_id, country_id);
  COMMIT;


  CREATE TABLE person(
     id serial constraint person_pk primary key,
     name varchar(60) not null unique,
     photo varchar(255) not null,
     description varchar(1024) not null
  );

  CREATE TABLE IF NOT EXISTS position(
      id serial constraint position_pk primary key,
      name varchar(32) not null unique
  );

  BEGIN;
  CREATE TABLE movies_staff(
       id serial constraint movies_staff_pk primary key,
       movie_id int,
       person_id int,
       position_id int
  );

  create unique index movies_staff_uindex
      on movies_staff (movie_id, person_id, position_id);
  COMMIT;

EOSQL
