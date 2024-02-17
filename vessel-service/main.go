package main

import (
	"fmt"
	grpcMicro "go.unistack.org/micro-server-grpc/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/server"
	"log"
	"os"
	pb "vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Server(grpcMicro.NewServer(server.Address(":50051"))),
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	err = pb.RegisterVesselServiceServer(srv.Server(), &handler{session})

	if err != nil {
		log.Fatalf("failed to register VesselServiceServer: %v", err)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
