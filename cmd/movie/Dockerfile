FROM golang:1.18 AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .env .
COPY . .

RUN CGO_ENABLED=0 go build ./cmd/movie/movie.go

FROM alpine

WORKDIR /app

COPY --from=build /build/movie .
COPY --from=build /build/.env .

RUN chmod +x ./movie

EXPOSE 5001/tcp

ENTRYPOINT ["./movie"]