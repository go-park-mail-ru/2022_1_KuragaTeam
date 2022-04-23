package main

import (
	"log"
	"myapp/internal/composites"
	"myapp/internal/microservices/authorization/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		log.Fatal("postgres db composite failed")
	}

	redisComposite, err := composites.NewRedisComposite()
	if err != nil {
		log.Fatal("redis composite failed")
	}

	authComposite, err := composites.NewAuthComposite(postgresDBC, redisComposite)
	if err != nil {
		log.Fatal("user composite failed")
	}

	listen, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Println("CANNOT LISTEN PORT: ", "localhost:5555", err.Error())
	}

	server := grpc.NewServer()

	proto.RegisterAuthorizationServer(server, authComposite.Service)
	log.Printf("STARTED AUTHORIZATION MICROSERVICE ON %s", "localhost:5555")
	err = server.Serve(listen)
	if err != nil {
		log.Println("CANNOT LISTEN PORT: ", "localhost:5555", err.Error())
	}
}
