// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dps

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

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	GetFirst(ctx context.Context, in *GetFirstRequest, opts ...grpc.CallOption) (*GetFirstResponse, error)
	GetLast(ctx context.Context, in *GetLastRequest, opts ...grpc.CallOption) (*GetLastResponse, error)
	GetHeader(ctx context.Context, in *GetHeaderRequest, opts ...grpc.CallOption) (*GetHeaderResponse, error)
	GetCommit(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*GetCommitResponse, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	GetRegisters(ctx context.Context, in *GetRegistersRequest, opts ...grpc.CallOption) (*GetRegistersResponse, error)
	GetHeight(ctx context.Context, in *GetHeightRequest, opts ...grpc.CallOption) (*GetHeightResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	ListTransactionsForBlock(ctx context.Context, in *ListTransactionsForBlockRequest, opts ...grpc.CallOption) (*ListTransactionsForBlockResponse, error)
	ListTransactionsForCollection(ctx context.Context, in *ListTransactionsForCollectionRequest, opts ...grpc.CallOption) (*ListTransactionsForCollectionResponse, error)
	ListCollectionsForBlock(ctx context.Context, in *ListCollectionsForBlockRequest, opts ...grpc.CallOption) (*ListCollectionsForBlockResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetFirst(ctx context.Context, in *GetFirstRequest, opts ...grpc.CallOption) (*GetFirstResponse, error) {
	out := new(GetFirstResponse)
	err := c.cc.Invoke(ctx, "/API/GetFirst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetLast(ctx context.Context, in *GetLastRequest, opts ...grpc.CallOption) (*GetLastResponse, error) {
	out := new(GetLastResponse)
	err := c.cc.Invoke(ctx, "/API/GetLast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetHeader(ctx context.Context, in *GetHeaderRequest, opts ...grpc.CallOption) (*GetHeaderResponse, error) {
	out := new(GetHeaderResponse)
	err := c.cc.Invoke(ctx, "/API/GetHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetCommit(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*GetCommitResponse, error) {
	out := new(GetCommitResponse)
	err := c.cc.Invoke(ctx, "/API/GetCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/API/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetRegisters(ctx context.Context, in *GetRegistersRequest, opts ...grpc.CallOption) (*GetRegistersResponse, error) {
	out := new(GetRegistersResponse)
	err := c.cc.Invoke(ctx, "/API/GetRegisters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetHeight(ctx context.Context, in *GetHeightRequest, opts ...grpc.CallOption) (*GetHeightResponse, error) {
	out := new(GetHeightResponse)
	err := c.cc.Invoke(ctx, "/API/GetHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/API/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTransactionsForBlock(ctx context.Context, in *ListTransactionsForBlockRequest, opts ...grpc.CallOption) (*ListTransactionsForBlockResponse, error) {
	out := new(ListTransactionsForBlockResponse)
	err := c.cc.Invoke(ctx, "/API/ListTransactionsForBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTransactionsForCollection(ctx context.Context, in *ListTransactionsForCollectionRequest, opts ...grpc.CallOption) (*ListTransactionsForCollectionResponse, error) {
	out := new(ListTransactionsForCollectionResponse)
	err := c.cc.Invoke(ctx, "/API/ListTransactionsForCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCollectionsForBlock(ctx context.Context, in *ListCollectionsForBlockRequest, opts ...grpc.CallOption) (*ListCollectionsForBlockResponse, error) {
	out := new(ListCollectionsForBlockResponse)
	err := c.cc.Invoke(ctx, "/API/ListCollectionsForBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations should embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	GetFirst(context.Context, *GetFirstRequest) (*GetFirstResponse, error)
	GetLast(context.Context, *GetLastRequest) (*GetLastResponse, error)
	GetHeader(context.Context, *GetHeaderRequest) (*GetHeaderResponse, error)
	GetCommit(context.Context, *GetCommitRequest) (*GetCommitResponse, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	GetRegisters(context.Context, *GetRegistersRequest) (*GetRegistersResponse, error)
	GetHeight(context.Context, *GetHeightRequest) (*GetHeightResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	ListTransactionsForBlock(context.Context, *ListTransactionsForBlockRequest) (*ListTransactionsForBlockResponse, error)
	ListTransactionsForCollection(context.Context, *ListTransactionsForCollectionRequest) (*ListTransactionsForCollectionResponse, error)
	ListCollectionsForBlock(context.Context, *ListCollectionsForBlockRequest) (*ListCollectionsForBlockResponse, error)
}

// UnimplementedAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) GetFirst(context.Context, *GetFirstRequest) (*GetFirstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirst not implemented")
}
func (UnimplementedAPIServer) GetLast(context.Context, *GetLastRequest) (*GetLastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLast not implemented")
}
func (UnimplementedAPIServer) GetHeader(context.Context, *GetHeaderRequest) (*GetHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeader not implemented")
}
func (UnimplementedAPIServer) GetCommit(context.Context, *GetCommitRequest) (*GetCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommit not implemented")
}
func (UnimplementedAPIServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedAPIServer) GetRegisters(context.Context, *GetRegistersRequest) (*GetRegistersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisters not implemented")
}
func (UnimplementedAPIServer) GetHeight(context.Context, *GetHeightRequest) (*GetHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeight not implemented")
}
func (UnimplementedAPIServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedAPIServer) ListTransactionsForBlock(context.Context, *ListTransactionsForBlockRequest) (*ListTransactionsForBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsForBlock not implemented")
}
func (UnimplementedAPIServer) ListTransactionsForCollection(context.Context, *ListTransactionsForCollectionRequest) (*ListTransactionsForCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsForCollection not implemented")
}
func (UnimplementedAPIServer) ListCollectionsForBlock(context.Context, *ListCollectionsForBlockRequest) (*ListCollectionsForBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollectionsForBlock not implemented")
}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_GetFirst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFirstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetFirst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetFirst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetFirst(ctx, req.(*GetFirstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetLast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetLast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetLast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetLast(ctx, req.(*GetLastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetHeader(ctx, req.(*GetHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetCommit(ctx, req.(*GetCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetRegisters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetRegisters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetRegisters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetRegisters(ctx, req.(*GetRegistersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetHeight(ctx, req.(*GetHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTransactionsForBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsForBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListTransactionsForBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListTransactionsForBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListTransactionsForBlock(ctx, req.(*ListTransactionsForBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTransactionsForCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsForCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListTransactionsForCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListTransactionsForCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListTransactionsForCollection(ctx, req.(*ListTransactionsForCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListCollectionsForBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsForBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListCollectionsForBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListCollectionsForBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListCollectionsForBlock(ctx, req.(*ListCollectionsForBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFirst",
			Handler:    _API_GetFirst_Handler,
		},
		{
			MethodName: "GetLast",
			Handler:    _API_GetLast_Handler,
		},
		{
			MethodName: "GetHeader",
			Handler:    _API_GetHeader_Handler,
		},
		{
			MethodName: "GetCommit",
			Handler:    _API_GetCommit_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _API_GetEvents_Handler,
		},
		{
			MethodName: "GetRegisters",
			Handler:    _API_GetRegisters_Handler,
		},
		{
			MethodName: "GetHeight",
			Handler:    _API_GetHeight_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _API_GetTransaction_Handler,
		},
		{
			MethodName: "ListTransactionsForBlock",
			Handler:    _API_ListTransactionsForBlock_Handler,
		},
		{
			MethodName: "ListTransactionsForCollection",
			Handler:    _API_ListTransactionsForCollection_Handler,
		},
		{
			MethodName: "ListCollectionsForBlock",
			Handler:    _API_ListCollectionsForBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
