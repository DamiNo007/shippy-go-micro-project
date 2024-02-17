// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// versions:
// - protoc-gen-go-micro v3.10.4
// - protoc              v3.6.1
// source: consignment.proto

package consignment

import (
	context "context"
	client "go.unistack.org/micro/v3/client"
)

var (
	ShippingServiceName = "ShippingService"
)

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, req *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, req *GetRequest, opts ...client.CallOption) (*Response, error)
}

type ShippingServiceServer interface {
	CreateConsignment(ctx context.Context, req *Consignment, rsp *Response) error
	GetConsignments(ctx context.Context, req *GetRequest, rsp *Response) error
}