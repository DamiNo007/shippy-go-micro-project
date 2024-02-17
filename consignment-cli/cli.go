package main

import (
	"encoding/json"
	grpcClientMicro "go.unistack.org/micro-client-grpc/v3"
	protoCodec "go.unistack.org/micro-codec-grpc/v3"
	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/metadata"
	"log"
	"os"
	"time"

	pb "github.com/DamiNo007/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("Did not connect: %v", err)
	//}
	//defer conn.Close()

	clt := pb.NewShippingServiceClient(
		"consignment-service:50051",
		grpcClientMicro.NewClient(
			client.Codec("application/grpc", protoCodec.NewCodec()),
		),
	)

	// Contact the server and print out its response.
	file := defaultFilename

	var token string

	log.Println(os.Args)

	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	token = os.Getenv("TOKEN")

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewOutgoingContext(context.Background(), map[string]string{
		"token": token,
	})

	for i := 0; i <= 100; i++ {
		r, err := clt.CreateConsignment(ctx, consignment)
		if err != nil {
			log.Fatalf("Could not greet: %v", err)
		}
		log.Printf("Created: %t", r.Created)

		getAll, err := clt.GetConsignments(ctx, &pb.GetRequest{})
		if err != nil {
			log.Fatalf("Could not list consignments: %v", err)
		}
		for _, v := range getAll.Consignments {
			log.Println(v)
		}

		time.Sleep(10 * time.Second)
	}
}
