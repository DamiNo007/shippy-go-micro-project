FROM golang:1.20.13 as builder

WORKDIR /go/src/shippy/user-service

COPY ./user-service .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/shippy/user-service/user-service .

CMD ["./user-service"]