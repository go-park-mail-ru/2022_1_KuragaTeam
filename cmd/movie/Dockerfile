#Build
FROM golang:1.18 AS build

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY ./. .

RUN go build ./cmd/movie/movie.go

#Environment
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/movie .

EXPOSE 5001/tcp

CMD ["./movie"]