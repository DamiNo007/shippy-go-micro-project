package main

import (
	pb "consignment-service/proto/consignment"
	"context"
	"errors"
	"fmt"
	userPb "github.com/DamiNo007/user-service/proto/user"
	vesselPb "github.com/DamiNo007/vessel-service/proto/vessel"
	grpcClientMicro "go.unistack.org/micro-client-grpc/v3"
	protoCodec "go.unistack.org/micro-codec-grpc/v3"
	grpcServerMicro "go.unistack.org/micro-server-grpc/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/metadata"
	"go.unistack.org/micro/v3/server"
	"log"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	//lis, err := net.Listen("tcp", port)
	//
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//s := grpc.NewServer()
	//
	//pb.RegisterShippingServiceServer(s, &service{repo})
	//
	//reflection.Register(s)
	//
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}

	srv := micro.NewService(
		micro.Server(
			grpcServerMicro.NewServer(
				server.Address(":50051"),
				server.WrapHandler(AuthWrapper),
			),
		),
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselPb.NewVesselServiceClient(
		"vessel-service:50051",
		grpcClientMicro.NewClient(
			client.Codec("application/grpc", protoCodec.NewCodec()),
		),
	)

	srv.Init()

	err = pb.RegisterShippingServiceServer(srv.Server(), &handler{session, vesselClient})

	if err != nil {
		log.Fatalf("failed to register ShippingServiceServer: %v", err)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp any) error {
		fmt.Printf("Type of context: %T\n", ctx)
		meta, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userPb.NewUserServiceClient(
			"user-service:50051",
			grpcClientMicro.NewClient(
				client.Codec("application/grpc", protoCodec.NewCodec()),
			),
		)

		authResp, err := authClient.ValidateToken(ctx, &userPb.Token{
			Token: token,
		})
		log.Println("Auth resp:", authResp)
		log.Println("Err:", err)
		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
