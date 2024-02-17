PROTO_PATH = ./proto/consignment

create-network:
	docker network create shipping-project-network

install-dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install go.unistack.org/protoc-gen-go-micro/v3

generate:
	#protoc --go_out=. --go_opt=paths=source_relative \
#                      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#                      proto/consignment/consignment.proto
	protoc -I$(PROTO_PATH) \
          --go-micro_out=components="micro|grpc",debug=true,tag_path=$(PROTO_PATH)/*.proto,paths=source_relative:$(PROTO_PATH) \
          $(PROTO_PATH)/*.proto
	protoc -I$(PROTO_PATH) \
          --go_out=paths=source_relative:$(PROTO_PATH) $(PROTO_PATH)/*.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 --network shipping-project-network --name consignment-service consignment-service

start:
	docker start consignment-service
	docker logs consignment-service -f

stop:
	docker stop consignment-service


delete-container:
	docker rm consignment-service
	docker rmi consignment-service