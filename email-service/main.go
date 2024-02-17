package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/DamiNo007/user-service/proto/user"
	_ "go.unistack.org/micro-broker-nats/v3"
	nats "go.unistack.org/micro-broker-nats/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/broker"
	"go.unistack.org/micro/v3/server"
	"log"
	"os"
)

const onUserCreatedTopic = "user.created"

func sendEmail(user *pb.User) error {
	log.Println("Sending email to:", user.Name)
	return nil
}

func onUserCreatedHandler(p broker.Event) error {
	var user *pb.User

	if err := json.Unmarshal(p.Message().Body, &user); err != nil {
		return err
	}

	log.Println(user)

	go sendEmail(user)

	return nil
}

func main() {
	broker := nats.NewBroker(
		broker.Addrs(os.Getenv("BROKER_HOST")),
	)

	srv := micro.NewService(
		micro.Server(
			server.NewServer(
				server.Address(":50051"),
				server.Broker(broker),
			),
		),
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	ctx := micro.NewContext(context.Background(), srv)

	pubSub := srv.Server().Options().Broker

	if err := pubSub.Connect(ctx); err != nil {
		fmt.Printf("could not connect to broker: %v", err.Error())
		log.Fatal(err)
	}

	_, err := pubSub.Subscribe(ctx, onUserCreatedTopic, onUserCreatedHandler)

	if err != nil {
		log.Println(err)
	}

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
