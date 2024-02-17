FROM golang:1.20.13 as builder

WORKDIR /go/src/shippy/user-cli

COPY ./user-cli .

ADD ./user-service ../user-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/shippy/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]