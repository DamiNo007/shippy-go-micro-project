package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/DamiNo007/user-service/proto/user"
	grpcClientMicro "go.unistack.org/micro-client-grpc/v3"
	protoCodec "go.unistack.org/micro-codec-grpc/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/client"
	"log"
)

var (
	name     string
	email    string
	password string
	company  string
)

func init() {
	flag.StringVar(&name, "name", "John", "Your name")
	flag.StringVar(&email, "email", "john007@gmail.com", "Your email")
	flag.StringVar(&password, "password", "John@99@12345", "Your password")
	flag.StringVar(&company, "company", "Google", "Your company")
}

type preStartFunc func(ctx context.Context) error

func preStart(client pb.UserServiceClient) preStartFunc {
	return func(ctx context.Context) error {
		r, err := client.Create(context.Background(), &pb.User{
			Name:     name,
			Email:    email,
			Password: password,
			Company:  company,
		})

		if err != nil {
			log.Printf("Could not create User: %v", err.Error())
			return err
		}

		log.Printf("Created: %v", r.User.Id)

		getAll, err := client.GetAll(context.Background(), &pb.Request{})

		if err != nil {
			log.Fatalf("Could not list users: %v", err)
		}

		for _, v := range getAll.Users {
			log.Println(v)
		}

		authResponse, err := client.Auth(context.TODO(), &pb.User{
			Email:    email,
			Password: password,
		})

		if err != nil {
			log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
		}

		log.Printf("Your access token is: %s \n", authResponse.Token)

		return nil
	}
}

func main() {
	flag.Parse()

	fmt.Println("name: ", name)

	clt := pb.NewUserServiceClient(
		"user-service:50051",
		grpcClientMicro.NewClient(
			client.Codec("application/grpc", protoCodec.NewCodec()),
		),
	)

	service := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	if err := service.Init(
		micro.BeforeStart(preStart(clt)),
	); err != nil {
		log.Fatalf("Could not create: %v", err)
	}

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
