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

RUN CGO_ENABLED=0 go build ./cmd/main/main.go

##
## Deploy
##

FROM alpine

WORKDIR /app

COPY --from=build /build/main .
COPY --from=build /build/.env .
COPY --from=build /build/default_avatar.webp .

RUN chmod +x ./main

EXPOSE 1323/tcp

ENTRYPOINT ["./main"]