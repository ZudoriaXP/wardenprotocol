// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: warden/act/v1beta1/query.proto

package actv1beta1

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

const (
	Query_Params_FullMethodName           = "/warden.act.v1beta1.Query/Params"
	Query_Actions_FullMethodName          = "/warden.act.v1beta1.Query/Actions"
	Query_Rules_FullMethodName            = "/warden.act.v1beta1.Query/Rules"
	Query_SimulateRule_FullMethodName     = "/warden.act.v1beta1.Query/SimulateRule"
	Query_RuleById_FullMethodName         = "/warden.act.v1beta1.Query/RuleById"
	Query_ActionsByAddress_FullMethodName = "/warden.act.v1beta1.Query/ActionsByAddress"
	Query_ActionById_FullMethodName       = "/warden.act.v1beta1.Query/ActionById"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Actions items.
	Actions(ctx context.Context, in *QueryActionsRequest, opts ...grpc.CallOption) (*QueryActionsResponse, error)
	// Queries a list of Rules items.
	Rules(ctx context.Context, in *QueryRulesRequest, opts ...grpc.CallOption) (*QueryRulesResponse, error)
	// Queries to simulate a Rule
	SimulateRule(ctx context.Context, in *QuerySimulateRuleRequest, opts ...grpc.CallOption) (*QuerySimulateRuleResponse, error)
	// Queries a list of RuleById items.
	RuleById(ctx context.Context, in *QueryRuleByIdRequest, opts ...grpc.CallOption) (*QueryRuleByIdResponse, error)
	// Queries a list of Actions items by one participant address.
	ActionsByAddress(ctx context.Context, in *QueryActionsByAddressRequest, opts ...grpc.CallOption) (*QueryActionsByAddressResponse, error)
	ActionById(ctx context.Context, in *QueryActionByIdRequest, opts ...grpc.CallOption) (*QueryActionByIdResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Actions(ctx context.Context, in *QueryActionsRequest, opts ...grpc.CallOption) (*QueryActionsResponse, error) {
	out := new(QueryActionsResponse)
	err := c.cc.Invoke(ctx, Query_Actions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Rules(ctx context.Context, in *QueryRulesRequest, opts ...grpc.CallOption) (*QueryRulesResponse, error) {
	out := new(QueryRulesResponse)
	err := c.cc.Invoke(ctx, Query_Rules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SimulateRule(ctx context.Context, in *QuerySimulateRuleRequest, opts ...grpc.CallOption) (*QuerySimulateRuleResponse, error) {
	out := new(QuerySimulateRuleResponse)
	err := c.cc.Invoke(ctx, Query_SimulateRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RuleById(ctx context.Context, in *QueryRuleByIdRequest, opts ...grpc.CallOption) (*QueryRuleByIdResponse, error) {
	out := new(QueryRuleByIdResponse)
	err := c.cc.Invoke(ctx, Query_RuleById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ActionsByAddress(ctx context.Context, in *QueryActionsByAddressRequest, opts ...grpc.CallOption) (*QueryActionsByAddressResponse, error) {
	out := new(QueryActionsByAddressResponse)
	err := c.cc.Invoke(ctx, Query_ActionsByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ActionById(ctx context.Context, in *QueryActionByIdRequest, opts ...grpc.CallOption) (*QueryActionByIdResponse, error) {
	out := new(QueryActionByIdResponse)
	err := c.cc.Invoke(ctx, Query_ActionById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Actions items.
	Actions(context.Context, *QueryActionsRequest) (*QueryActionsResponse, error)
	// Queries a list of Rules items.
	Rules(context.Context, *QueryRulesRequest) (*QueryRulesResponse, error)
	// Queries to simulate a Rule
	SimulateRule(context.Context, *QuerySimulateRuleRequest) (*QuerySimulateRuleResponse, error)
	// Queries a list of RuleById items.
	RuleById(context.Context, *QueryRuleByIdRequest) (*QueryRuleByIdResponse, error)
	// Queries a list of Actions items by one participant address.
	ActionsByAddress(context.Context, *QueryActionsByAddressRequest) (*QueryActionsByAddressResponse, error)
	ActionById(context.Context, *QueryActionByIdRequest) (*QueryActionByIdResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Actions(context.Context, *QueryActionsRequest) (*QueryActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Actions not implemented")
}
func (UnimplementedQueryServer) Rules(context.Context, *QueryRulesRequest) (*QueryRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rules not implemented")
}
func (UnimplementedQueryServer) SimulateRule(context.Context, *QuerySimulateRuleRequest) (*QuerySimulateRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimulateRule not implemented")
}
func (UnimplementedQueryServer) RuleById(context.Context, *QueryRuleByIdRequest) (*QueryRuleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RuleById not implemented")
}
func (UnimplementedQueryServer) ActionsByAddress(context.Context, *QueryActionsByAddressRequest) (*QueryActionsByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionsByAddress not implemented")
}
func (UnimplementedQueryServer) ActionById(context.Context, *QueryActionByIdRequest) (*QueryActionByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionById not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Actions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Actions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Actions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Actions(ctx, req.(*QueryActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Rules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Rules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Rules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Rules(ctx, req.(*QueryRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SimulateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySimulateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SimulateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SimulateRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SimulateRule(ctx, req.(*QuerySimulateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RuleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRuleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RuleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_RuleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RuleById(ctx, req.(*QueryRuleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ActionsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ActionsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ActionsByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ActionsByAddress(ctx, req.(*QueryActionsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ActionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ActionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ActionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ActionById(ctx, req.(*QueryActionByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warden.act.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Actions",
			Handler:    _Query_Actions_Handler,
		},
		{
			MethodName: "Rules",
			Handler:    _Query_Rules_Handler,
		},
		{
			MethodName: "SimulateRule",
			Handler:    _Query_SimulateRule_Handler,
		},
		{
			MethodName: "RuleById",
			Handler:    _Query_RuleById_Handler,
		},
		{
			MethodName: "ActionsByAddress",
			Handler:    _Query_ActionsByAddress_Handler,
		},
		{
			MethodName: "ActionById",
			Handler:    _Query_ActionById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warden/act/v1beta1/query.proto",
}
