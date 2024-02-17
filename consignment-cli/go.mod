module consignment-cli

go 1.20

replace (
	github.com/DamiNo007/consignment-service => ../consignment-service
	github.com/DamiNo007/vessel-service => ../vessel-service
	github.com/DamiNo007/user-service => ../user-service
)

require (
	github.com/DamiNo007/consignment-service v0.0.0-00010101000000-000000000000
	go.unistack.org/micro-client-grpc/v3 v3.11.3
	go.unistack.org/micro-codec-grpc/v3 v3.10.0
	go.unistack.org/micro/v3 v3.10.38
	golang.org/x/net v0.21.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	go.unistack.org/micro-proto/v3 v3.3.1 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
