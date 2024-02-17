build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run --network shipping-project-network --name consignment-cli consignment-cli

start:
	docker start consignment-cli
	docker logs consignment-cli -f

stop:
	docker stop consignment-cli

delete-container:
	docker rm consignment-cli
	docker rmi consignment-cli