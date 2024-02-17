package main

import (
	"context"
	"fmt"
	_ "go.unistack.org/micro-broker-nats/v3"
	nats "go.unistack.org/micro-broker-nats/v3"
	grpcMicro "go.unistack.org/micro-server-grpc/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/broker"
	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/server"
	"log"
	"os"
	pb "user-service/proto/user"
)

func main() {
	db, err := CreateConnection()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	natsBroker := nats.NewBroker(
		broker.Addrs(os.Getenv("BROKER_HOST")),
	)

	srv := micro.NewService(
		micro.Server(grpcMicro.NewServer(
			server.Address(":50051"),
		)),
		micro.Client(client.NewClient(
			client.Broker(natsBroker),
		)),
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	ctx := micro.NewContext(context.Background(), srv)

	pubSub := srv.Client().Options().Broker

	if err := pubSub.Connect(ctx); err != nil {
		fmt.Printf("could not connect to broker: %v", err.Error())
		log.Fatal(err)
	}

	err = pb.RegisterUserServiceServer(srv.Server(), &handler{repo, tokenService, pubSub})

	if err != nil {
		log.Fatalf("failed to register VesselServiceServer: %v", err)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
