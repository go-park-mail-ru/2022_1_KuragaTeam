##
## Build
##

FROM golang:1.18 AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/authorization/authorization.go

##
## Deploy
##

FROM alpine

WORKDIR /app

COPY --from=build /build/authorization .
COPY --from=build /build/.env .

RUN chmod +x ./authorization

EXPOSE 5555/tcp

ENTRYPOINT ["./authorization"]