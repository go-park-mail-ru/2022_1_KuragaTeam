#!/bin/bash

moq -out internal/microservices/movie/usecase/movie_usecase_mock.go -pkg usecase ./internal/microservices/movie/proto MoviesServer:MockMoviesServer
moq -out internal/microservices/movie/repository/movie_repository_mock.go -pkg repository ./internal/microservices/movie Storage:MockMovieStorage

moq -out internal/genre/repository/genre_repository_mock.go -pkg repository ./internal/genre Storage:MockGenreStorage
moq -out internal/country/repository/country_repository_mock.go -pkg repository ./internal/country Storage:MockCountryStorage

moq -out internal/persons/usecase/persons_usecase_mock.go -pkg usecase ./internal/persons Service:MockPersonsService
moq -out internal/persons/repository/persons_repository_mock.go -pkg repository ./internal/persons Storage:MockPersonsStorage

moq -out mock/position_repository_mock.go -pkg mock ./internal/position Storage:MockPositionStorage

moq -out internal/microservices/compilations/usecase/compilations_usecase_mock.go -pkg usecase ./internal/microservices/compilations/proto MovieCompilationsServer:MockMovieCompilationsServer
moq -out internal/microservices/compilations/repository/compilations_repository_mock.go -pkg repository ./internal/microservices/compilations Storage:MockMovieCompilationStorage

mockgen -source=./internal/microservices/authorization/repository.go -destination=./internal/microservices/authorization/repository/auth_repository_mock.go -package=repository
mockgen -source=./internal/microservices/profile/repository.go -destination=./internal/microservices/profile/repository/profile_repository_mock.go -package=repository
mockgen -source=./internal/microservices/authorization/proto/authorization_grpc.pb.go -destination=./internal/microservices/authorization/usecase/auth_usecase_mock.go -package=usecase
mockgen -source=./internal/microservices/profile/proto/profile_grpc.pb.go -destination=./internal/microservices/profile/usecase/profile_usecase_mock.go -package=usecase



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