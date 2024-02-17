// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// versions:
// - protoc-gen-go-micro v3.10.4
// - protoc              v3.6.1
// source: vessel.proto

package vessel

import (
	context "context"
	client "go.unistack.org/micro/v3/client"
)

var (
	VesselServiceName = "VesselService"
)

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, req *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, req *Vessel, opts ...client.CallOption) (*Response, error)
}

type VesselServiceServer interface {
	FindAvailable(ctx context.Context, req *Specification, rsp *Response) error
	Create(ctx context.Context, req *Vessel, rsp *Response) error
}