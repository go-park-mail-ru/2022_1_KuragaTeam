package main

import (
	"log"
	"myapp/internal/composites"
	"myapp/internal/microservices/profile/proto"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		log.Fatal("postgres db composite failed")
	}

	minioComposite, err := composites.NewMinioComposite()
	if err != nil {
		log.Fatal("minio composite failed")
	}

	profileComposite, err := composites.NewProfileComposite(postgresDBC, minioComposite)
	if err != nil {
		log.Fatal("profile composite failed")
	}

	listen, err := net.Listen("tcp", os.Getenv("PROFILE_MICROSERVICE_HOST")+":"+
		os.Getenv("PROFILE_MICROSERVICE_PORT"))
	if err != nil {
		log.Fatal("CANNOT LISTEN PORT: ", os.Getenv("PROFILE_MICROSERVICE_HOST")+":"+
			os.Getenv("PROFILE_MICROSERVICE_PORT"), err.Error())
	}

	server := grpc.NewServer()

	proto.RegisterProfileServer(server, profileComposite.Service)
	log.Printf("STARTED AUTHORIZATION MICROSERVICE ON %s", os.Getenv("PROFILE_MICROSERVICE_HOST")+":"+
		os.Getenv("PROFILE_MICROSERVICE_PORT"))
	err = server.Serve(listen)
	if err != nil {
		log.Println("CANNOT LISTEN PORT: ", os.Getenv("PROFILE_MICROSERVICE_HOST")+":"+
			os.Getenv("PROFILE_MICROSERVICE_PORT"), err.Error())
	}
}
