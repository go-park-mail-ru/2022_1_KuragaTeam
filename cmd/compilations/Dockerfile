FROM golang:1.18 AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .env .
COPY . .

RUN CGO_ENABLED=0 go build ./cmd/compilations/compilations.go

FROM alpine

WORKDIR /app

COPY --from=build /build/compilations .
COPY --from=build /build/.env .

RUN chmod +x ./compilations

EXPOSE 5002/tcp

ENTRYPOINT ["./compilations"]