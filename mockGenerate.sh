#!/bin/bash

moq -out mock/movie_usecase_mock.go -pkg mock ./internal/movie Service:MockMovieService
moq -out mock/movie_repository_mock.go -pkg mock ./internal/movie Storage:MockMovieStorage

moq -out mock/genre_repository_mock.go -pkg mock ./internal/genre Storage:MockGenreStorage
moq -out mock/country_repository_mock.go -pkg mock ./internal/country Storage:MockCountryStorage

moq -out mock/persons_usecase_mock.go -pkg mock ./internal/persons Service:MockPersonsService
moq -out mock/persons_repository_mock.go -pkg mock ./internal/persons Storage:MockPersonsStorage

moq -out mock/position_repository_mock.go -pkg mock ./internal/position Storage:MockPositionStorage

moq -out mock/movieCompilations_usecase_mock.go -pkg mock ./internal/moviesCompilations Service:MockMovieCompilationService
moq -out mock/movieCompilations_repository_mock.go -pkg mock ./internal/moviesCompilations Storage:MockMovieCompilationStorage

mockgen -source=./internal/user/repository.go -destination=./mock/userRepositoryMock.go -package=mock
mockgen -source=./internal/user/usecase.go -destination=./mock/userUsecaseMock.go -package=mock

touch .env
printf "DBHOST=postgres
DBPORT=5432
DBUSER=docker
DBPASSWORD=docker
DBNAME=docker
REDISHOST=redis
REDISPORT=6379
REDISPROTOCOL=tcp
MINIOURL=minio:9000
NGINX=localhost:8000
MINIOUSER=minio
MINIOPASSWORD=minio123
CSRF_SECRET=secret\n" > .env

cp .env ./internal/movie/usecase
cp .env ./internal/moviesCompilations/usecase
cp .env ./internal/persons/usecase
cp .env ./internal/user/usecase
cp .env ./internal/user/repository
cp .env ./internal/utils/images

go test -coverpkg=./... -coverprofile cover.out.tmp ./...
cat cover.out.tmp | grep -v "monitoring" | grep -v "easyjson" | grep -v "mock_*" | grep -v ".pb.go" | grep -v ".pb" | grep -v "middleware.go" | grep -v "/cmd*"> cover.out
go tool cover -func cover.out