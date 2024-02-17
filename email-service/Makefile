PROTO_PATH = ./proto/user

create-network:
	docker network create shipping-project-network

install-dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install go.unistack.org/protoc-gen-go-micro/v3

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t email-service .

run:
	docker run -p 50051:50051 --network shipping-project-network --name email-service email-service

start:
	docker start email-service
	docker logs email-service -f

stop:
	docker stop email-service


delete-container:
	docker rm email-service
	docker rmi email-service