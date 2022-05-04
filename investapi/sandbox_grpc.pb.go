// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: tinkoff/invest/grpc/sandbox.proto

package investapi

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

// SandboxServiceClient is the client API for SandboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SandboxServiceClient interface {
	//Метод регистрации счёта в песочнице.
	OpenSandboxAccount(ctx context.Context, in *OpenSandboxAccountRequest, opts ...grpc.CallOption) (*OpenSandboxAccountResponse, error)
	//Метод получения счетов в песочнице.
	GetSandboxAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error)
	//Метод закрытия счёта в песочнице.
	CloseSandboxAccount(ctx context.Context, in *CloseSandboxAccountRequest, opts ...grpc.CallOption) (*CloseSandboxAccountResponse, error)
	//Метод выставления торгового поручения в песочнице.
	PostSandboxOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error)
	//Метод получения списка активных заявок по счёту в песочнице.
	GetSandboxOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	//Метод отмены торгового поручения в песочнице.
	CancelSandboxOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	//Метод получения статуса заявки в песочнице.
	GetSandboxOrderState(ctx context.Context, in *GetOrderStateRequest, opts ...grpc.CallOption) (*OrderState, error)
	//Метод получения позиций по виртуальному счёту песочницы.
	GetSandboxPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error)
	//Метод получения операций в песочнице по номеру счёта.
	GetSandboxOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error)
	//Метод получения портфолио в песочнице.
	GetSandboxPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error)
	//Метод пополнения счёта в песочнице.
	SandboxPayIn(ctx context.Context, in *SandboxPayInRequest, opts ...grpc.CallOption) (*SandboxPayInResponse, error)
}

type sandboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSandboxServiceClient(cc grpc.ClientConnInterface) SandboxServiceClient {
	return &sandboxServiceClient{cc}
}

func (c *sandboxServiceClient) OpenSandboxAccount(ctx context.Context, in *OpenSandboxAccountRequest, opts ...grpc.CallOption) (*OpenSandboxAccountResponse, error) {
	out := new(OpenSandboxAccountResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/OpenSandboxAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error) {
	out := new(GetAccountsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) CloseSandboxAccount(ctx context.Context, in *CloseSandboxAccountRequest, opts ...grpc.CallOption) (*CloseSandboxAccountResponse, error) {
	out := new(CloseSandboxAccountResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/CloseSandboxAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) PostSandboxOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error) {
	out := new(PostOrderResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) CancelSandboxOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOrderState(ctx context.Context, in *GetOrderStateRequest, opts ...grpc.CallOption) (*OrderState, error) {
	out := new(OrderState)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error) {
	out := new(PositionsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error) {
	out := new(OperationsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) GetSandboxPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error) {
	out := new(PortfolioResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxServiceClient) SandboxPayIn(ctx context.Context, in *SandboxPayInRequest, opts ...grpc.CallOption) (*SandboxPayInResponse, error) {
	out := new(SandboxPayInResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.SandboxService/SandboxPayIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SandboxServiceServer is the server API for SandboxService service.
// All implementations must embed UnimplementedSandboxServiceServer
// for forward compatibility
type SandboxServiceServer interface {
	//Метод регистрации счёта в песочнице.
	OpenSandboxAccount(context.Context, *OpenSandboxAccountRequest) (*OpenSandboxAccountResponse, error)
	//Метод получения счетов в песочнице.
	GetSandboxAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error)
	//Метод закрытия счёта в песочнице.
	CloseSandboxAccount(context.Context, *CloseSandboxAccountRequest) (*CloseSandboxAccountResponse, error)
	//Метод выставления торгового поручения в песочнице.
	PostSandboxOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error)
	//Метод получения списка активных заявок по счёту в песочнице.
	GetSandboxOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	//Метод отмены торгового поручения в песочнице.
	CancelSandboxOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	//Метод получения статуса заявки в песочнице.
	GetSandboxOrderState(context.Context, *GetOrderStateRequest) (*OrderState, error)
	//Метод получения позиций по виртуальному счёту песочницы.
	GetSandboxPositions(context.Context, *PositionsRequest) (*PositionsResponse, error)
	//Метод получения операций в песочнице по номеру счёта.
	GetSandboxOperations(context.Context, *OperationsRequest) (*OperationsResponse, error)
	//Метод получения портфолио в песочнице.
	GetSandboxPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error)
	//Метод пополнения счёта в песочнице.
	SandboxPayIn(context.Context, *SandboxPayInRequest) (*SandboxPayInResponse, error)
	mustEmbedUnimplementedSandboxServiceServer()
}

// UnimplementedSandboxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSandboxServiceServer struct {
}

func (UnimplementedSandboxServiceServer) OpenSandboxAccount(context.Context, *OpenSandboxAccountRequest) (*OpenSandboxAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSandboxAccount not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxAccounts not implemented")
}
func (UnimplementedSandboxServiceServer) CloseSandboxAccount(context.Context, *CloseSandboxAccountRequest) (*CloseSandboxAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSandboxAccount not implemented")
}
func (UnimplementedSandboxServiceServer) PostSandboxOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSandboxOrder not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOrders not implemented")
}
func (UnimplementedSandboxServiceServer) CancelSandboxOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSandboxOrder not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOrderState(context.Context, *GetOrderStateRequest) (*OrderState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOrderState not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxPositions(context.Context, *PositionsRequest) (*PositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxPositions not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxOperations(context.Context, *OperationsRequest) (*OperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxOperations not implemented")
}
func (UnimplementedSandboxServiceServer) GetSandboxPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSandboxPortfolio not implemented")
}
func (UnimplementedSandboxServiceServer) SandboxPayIn(context.Context, *SandboxPayInRequest) (*SandboxPayInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SandboxPayIn not implemented")
}
func (UnimplementedSandboxServiceServer) mustEmbedUnimplementedSandboxServiceServer() {}

// UnsafeSandboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SandboxServiceServer will
// result in compilation errors.
type UnsafeSandboxServiceServer interface {
	mustEmbedUnimplementedSandboxServiceServer()
}

func RegisterSandboxServiceServer(s grpc.ServiceRegistrar, srv SandboxServiceServer) {
	s.RegisterService(&SandboxService_ServiceDesc, srv)
}

func _SandboxService_OpenSandboxAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSandboxAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).OpenSandboxAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/OpenSandboxAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).OpenSandboxAccount(ctx, req.(*OpenSandboxAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_CloseSandboxAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSandboxAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).CloseSandboxAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/CloseSandboxAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).CloseSandboxAccount(ctx, req.(*CloseSandboxAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_PostSandboxOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).PostSandboxOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/PostSandboxOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).PostSandboxOrder(ctx, req.(*PostOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_CancelSandboxOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).CancelSandboxOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/CancelSandboxOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).CancelSandboxOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOrderState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOrderState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOrderState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOrderState(ctx, req.(*GetOrderStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxPositions(ctx, req.(*PositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxOperations(ctx, req.(*OperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_GetSandboxPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).GetSandboxPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/GetSandboxPortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).GetSandboxPortfolio(ctx, req.(*PortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SandboxService_SandboxPayIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxPayInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServiceServer).SandboxPayIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.SandboxService/SandboxPayIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServiceServer).SandboxPayIn(ctx, req.(*SandboxPayInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SandboxService_ServiceDesc is the grpc.ServiceDesc for SandboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SandboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.SandboxService",
	HandlerType: (*SandboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenSandboxAccount",
			Handler:    _SandboxService_OpenSandboxAccount_Handler,
		},
		{
			MethodName: "GetSandboxAccounts",
			Handler:    _SandboxService_GetSandboxAccounts_Handler,
		},
		{
			MethodName: "CloseSandboxAccount",
			Handler:    _SandboxService_CloseSandboxAccount_Handler,
		},
		{
			MethodName: "PostSandboxOrder",
			Handler:    _SandboxService_PostSandboxOrder_Handler,
		},
		{
			MethodName: "GetSandboxOrders",
			Handler:    _SandboxService_GetSandboxOrders_Handler,
		},
		{
			MethodName: "CancelSandboxOrder",
			Handler:    _SandboxService_CancelSandboxOrder_Handler,
		},
		{
			MethodName: "GetSandboxOrderState",
			Handler:    _SandboxService_GetSandboxOrderState_Handler,
		},
		{
			MethodName: "GetSandboxPositions",
			Handler:    _SandboxService_GetSandboxPositions_Handler,
		},
		{
			MethodName: "GetSandboxOperations",
			Handler:    _SandboxService_GetSandboxOperations_Handler,
		},
		{
			MethodName: "GetSandboxPortfolio",
			Handler:    _SandboxService_GetSandboxPortfolio_Handler,
		},
		{
			MethodName: "SandboxPayIn",
			Handler:    _SandboxService_SandboxPayIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tinkoff/invest/grpc/sandbox.proto",
}
