build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-cli .

run:
	docker run --network shipping-project-network --name user-cli user-cli

start:
	docker start user-cli
	docker logs user-cli -f

stop:
	docker stop user-cli

delete-container:
	docker rm user-cli
	docker rmi user-cli