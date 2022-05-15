##
## Build
##

FROM golang:1.18 AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .env .
COPY . .

RUN CGO_ENABLED=0 go build ./cmd/profile/profile.go

##
## Deploy
##

FROM alpine

WORKDIR /app

COPY --from=build /build/profile .
COPY --from=build /build/.env .
COPY --from=build /build/default_avatar.webp .

RUN chmod +x ./profile

EXPOSE 6666/tcp

ENTRYPOINT ["./profile"]