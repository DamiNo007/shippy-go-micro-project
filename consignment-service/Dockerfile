FROM golang:1.20.13 as builder

WORKDIR /go/src/shippy/consignment-service

COPY ./consignment-service .

ADD ./vessel-service ../vessel-service
ADD ./user-service ../user-service

RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/shippy/consignment-service/consignment-service .

CMD ["./consignment-service"]