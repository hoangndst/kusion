// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: module.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Module_Generate_FullMethodName = "/Module/Generate"
)

// ModuleClient is the client API for Module service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModuleClient interface {
	Generate(ctx context.Context, in *GeneratorRequest, opts ...grpc.CallOption) (*GeneratorResponse, error)
}

type moduleClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleClient(cc grpc.ClientConnInterface) ModuleClient {
	return &moduleClient{cc}
}

func (c *moduleClient) Generate(ctx context.Context, in *GeneratorRequest, opts ...grpc.CallOption) (*GeneratorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneratorResponse)
	err := c.cc.Invoke(ctx, Module_Generate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModuleServer is the server API for Module service.
// All implementations must embed UnimplementedModuleServer
// for forward compatibility.
type ModuleServer interface {
	Generate(context.Context, *GeneratorRequest) (*GeneratorResponse, error)
	mustEmbedUnimplementedModuleServer()
}

// UnimplementedModuleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModuleServer struct{}

func (UnimplementedModuleServer) Generate(context.Context, *GeneratorRequest) (*GeneratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedModuleServer) mustEmbedUnimplementedModuleServer() {}
func (UnimplementedModuleServer) testEmbeddedByValue()                {}

// UnsafeModuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModuleServer will
// result in compilation errors.
type UnsafeModuleServer interface {
	mustEmbedUnimplementedModuleServer()
}

func RegisterModuleServer(s grpc.ServiceRegistrar, srv ModuleServer) {
	// If the following call pancis, it indicates UnimplementedModuleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Module_ServiceDesc, srv)
}

func _Module_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Module_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServer).Generate(ctx, req.(*GeneratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Module_ServiceDesc is the grpc.ServiceDesc for Module service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Module_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Module",
	HandlerType: (*ModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Module_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module.proto",
}
