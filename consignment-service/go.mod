module consignment-service

go 1.20

replace (
	github.com/DamiNo007/user-service => ../user-service
	github.com/DamiNo007/vessel-service => ../vessel-service
)

require (
	github.com/DamiNo007/user-service v0.0.0-00010101000000-000000000000
	github.com/DamiNo007/vessel-service v0.0.0-00010101000000-000000000000
	go.unistack.org/micro-client-grpc/v3 v3.11.3
	go.unistack.org/micro-codec-grpc/v3 v3.10.0
	go.unistack.org/micro-server-grpc/v3 v3.10.11
	go.unistack.org/micro/v3 v3.10.38
	google.golang.org/protobuf v1.32.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	go.unistack.org/micro-proto/v3 v3.3.1 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/grpc v1.61.0 // indirect
)
