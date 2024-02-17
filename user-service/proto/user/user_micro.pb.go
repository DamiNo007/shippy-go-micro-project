// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// versions:
// - protoc-gen-go-micro v3.10.4
// - protoc              v3.6.1
// source: user.proto

package user

import (
	context "context"
	client "go.unistack.org/micro/v3/client"
)

var (
	UserServiceName = "UserService"
)

type UserServiceClient interface {
	Create(ctx context.Context, req *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, req *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, req *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, req *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, req *Token, opts ...client.CallOption) (*Token, error)
}

type UserServiceServer interface {
	Create(ctx context.Context, req *User, rsp *Response) error
	Get(ctx context.Context, req *User, rsp *Response) error
	GetAll(ctx context.Context, req *Request, rsp *Response) error
	Auth(ctx context.Context, req *User, rsp *Token) error
	ValidateToken(ctx context.Context, req *Token, rsp *Token) error
}
