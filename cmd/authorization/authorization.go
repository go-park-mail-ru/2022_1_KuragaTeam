package main

import (
	"log"
	"myapp/internal/composites"
	"myapp/internal/microservices/authorization/proto"
	"net"
	"os"

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
		log.Fatal("auth composite failed")
	}

	listen, err := net.Listen("tcp", ":"+os.Getenv("AUTH_PORT"))
	if err != nil {
		log.Fatal("CANNOT LISTEN PORT: ", ":"+os.Getenv("AUTH_PORT"), err.Error())
	}

	server := grpc.NewServer()

	proto.RegisterAuthorizationServer(server, authComposite.Service)
	log.Printf("STARTED AUTHORIZATION MICROSERVICE ON %s", ":"+os.Getenv("AUTH_PORT"))
	err = server.Serve(listen)
	if err != nil {
		log.Println("CANNOT LISTEN PORT: ", ":"+os.Getenv("AUTH_PORT"), err.Error())
	}
}
