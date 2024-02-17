PROTO_PATH = ./proto/vessel

generate:
	protoc -I$(PROTO_PATH) \
		--go-micro_out=components="micro|grpc",debug=true,tag_path=$(PROTO_PATH)/*.proto,paths=source_relative:$(PROTO_PATH) \
		$(PROTO_PATH)/*.proto
	protoc -I$(PROTO_PATH) \
		--go_out=paths=source_relative:$(PROTO_PATH) $(PROTO_PATH)/*.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 --network shipping-project-network --name vessel-service vessel-service

start:
	docker start vessel-service
	docker logs vessel-service -f

stop:
	docker stop vessel-service

delete-container:
	docker rm vessel-service
	docker rmi vessel-service