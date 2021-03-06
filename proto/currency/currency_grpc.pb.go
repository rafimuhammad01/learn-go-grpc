// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package currency

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyClient interface {
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, "/currency.Currency/GetCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
// All implementations must embed UnimplementedCurrencyServer
// for forward compatibility
type CurrencyServer interface {
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	mustEmbedUnimplementedCurrencyServer()
}

// UnimplementedCurrencyServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (UnimplementedCurrencyServer) GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (UnimplementedCurrencyServer) mustEmbedUnimplementedCurrencyServer() {}

// UnsafeCurrencyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServer will
// result in compilation errors.
type UnsafeCurrencyServer interface {
	mustEmbedUnimplementedCurrencyServer()
}

func RegisterCurrencyServer(s grpc.ServiceRegistrar, srv CurrencyServer) {
	s.RegisterService(&Currency_ServiceDesc, srv)
}

func _Currency_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currency.Currency/GetCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Currency_ServiceDesc is the grpc.ServiceDesc for Currency service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Currency_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "currency.Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrency",
			Handler:    _Currency_GetCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}
