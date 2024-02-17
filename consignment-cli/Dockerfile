FROM golang:1.20.13 as builder

WORKDIR /go/src/shippy/consignment-cli

COPY ./consignment-cli .

ADD ./consignment-service ../consignment-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/shippy/consignment-cli/consignment.json .
COPY --from=builder /go/src/shippy/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]